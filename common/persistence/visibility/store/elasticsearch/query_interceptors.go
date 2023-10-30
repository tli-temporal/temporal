// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package elasticsearch

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	enumspb "go.temporal.io/api/enums/v1"

	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/common/persistence/visibility/store/query"
	"go.temporal.io/server/common/primitives/timestamp"
	"go.temporal.io/server/common/searchattribute"
)

type (
	nameInterceptor struct {
		namespace                      namespace.Name
		index                          string
		searchAttributesTypeMap        searchattribute.NameTypeMap
		searchAttributesMapperProvider searchattribute.MapperProvider
		seenNamespaceDivision          bool
	}

	valuesInterceptor struct {
		namespace                      namespace.Name
		searchAttributesTypeMap        searchattribute.NameTypeMap
		searchAttributesMapperProvider searchattribute.MapperProvider
	}
)

func newNameInterceptor(
	namespaceName namespace.Name,
	index string,
	saTypeMap searchattribute.NameTypeMap,
	searchAttributesMapperProvider searchattribute.MapperProvider,
) *nameInterceptor {
	return &nameInterceptor{
		namespace:                      namespaceName,
		index:                          index,
		searchAttributesTypeMap:        saTypeMap,
		searchAttributesMapperProvider: searchAttributesMapperProvider,
	}
}

func NewValuesInterceptor(
	namespaceName namespace.Name,
	saTypeMap searchattribute.NameTypeMap,
	searchAttributesMapperProvider searchattribute.MapperProvider,
) *valuesInterceptor {
	return &valuesInterceptor{
		namespace:                      namespaceName,
		searchAttributesTypeMap:        saTypeMap,
		searchAttributesMapperProvider: searchAttributesMapperProvider,
	}
}

func (ni *nameInterceptor) Name(name string, usage query.FieldNameUsage) (string, error) {
	fieldName := name
	if searchattribute.IsMappable(name) {
		mapper, err := ni.searchAttributesMapperProvider.GetMapper(ni.namespace)
		if err != nil {
			return "", err
		}
		if mapper != nil {
			fieldName, err = mapper.GetFieldName(name, ni.namespace.String())
			if err != nil {
				return "", err
			}
		}
	}

	fieldType, err := ni.searchAttributesTypeMap.GetType(fieldName)
	if err != nil {
		return "", query.NewConverterError("invalid search attribute: %s", name)
	}

	switch usage {
	case query.FieldNameFilter:
		if fieldName == searchattribute.TemporalNamespaceDivision {
			ni.seenNamespaceDivision = true
		}
	case query.FieldNameSorter:
		if fieldType == enumspb.INDEXED_VALUE_TYPE_TEXT {
			return "", query.NewConverterError(
				"unable to sort by field of %s type, use field of type %s",
				enumspb.INDEXED_VALUE_TYPE_TEXT.String(),
				enumspb.INDEXED_VALUE_TYPE_KEYWORD.String(),
			)
		}
	case query.FieldNameGroupBy:
		if fieldName != searchattribute.ExecutionStatus {
			return "", query.NewConverterError(
				"'group by' clause is only supported for %s search attribute",
				searchattribute.ExecutionStatus,
			)
		}
	}

	return fieldName, nil
}

func (vi *valuesInterceptor) Values(fieldName string, values ...interface{}) ([]interface{}, error) {
	fieldType, err := vi.searchAttributesTypeMap.GetType(fieldName)
	if err != nil {
		return nil, query.NewConverterError("invalid search attribute: %s", fieldName)
	}

	name := fieldName
	if searchattribute.IsMappable(fieldName) {
		mapper, err := vi.searchAttributesMapperProvider.GetMapper(vi.namespace)
		if err != nil {
			return nil, err
		}
		if mapper != nil {
			name, err = mapper.GetAlias(fieldName, vi.namespace.String())
			if err != nil {
				return nil, err
			}
		}
	}

	var result []interface{}
	for _, value := range values {
		value, err = vi.parseSystemSearchAttributeValues(fieldName, value)
		if err != nil {
			return nil, err
		}
		value, err = validateValueType(name, value, fieldType)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func (vi *valuesInterceptor) parseHHMMSSDuration(d string) (int64, error) {
	var hours, minutes, seconds, nanos int64
	_, err := fmt.Sscanf(d, "%d:%d:%d", &hours, &minutes, &seconds)
	if err != nil {
		return 0, errors.New("value is not a duration")
	}
	if hours < 0 {
		return 0, query.NewConverterError("invalid duration: hours must be positive number")
	}
	if minutes < 0 || minutes > 59 {
		return 0, query.NewConverterError("invalid duration: minutes must be from 0 to 59")
	}
	if seconds < 0 || seconds > 59 {
		return 0, query.NewConverterError("invalid duration: seconds must be from 0 to 59")
	}

	return hours*int64(time.Hour) + minutes*int64(time.Minute) + seconds*int64(time.Second) + nanos, nil
}

func (vi *valuesInterceptor) parseSystemSearchAttributeValues(name string, value any) (any, error) {
	switch name {
	case searchattribute.StartTime, searchattribute.CloseTime, searchattribute.ExecutionTime:
		if nanos, isNumber := value.(int64); isNumber {
			value = time.Unix(0, nanos).UTC().Format(time.RFC3339Nano)
		}
	case searchattribute.ExecutionStatus:
		if status, isNumber := value.(int64); isNumber {
			value = enumspb.WorkflowExecutionStatus_name[int32(status)]
		}
	case searchattribute.ExecutionDuration:
		if durationStr, isString := value.(string); isString {
			// To support durations passed as golang durations such as "300ms", "-1.5h" or "2h45m".
			// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
			// Custom timestamp.ParseDuration also supports "d" as additional unit for days.
			if duration, err := timestamp.ParseDuration(durationStr); err == nil {
				value = duration.Nanoseconds()
			} else {
				// To support "hh:mm:ss" durations.
				durationNanos, err := vi.parseHHMMSSDuration(durationStr)
				var converterErr *query.ConverterError
				if errors.As(err, &converterErr) {
					return nil, converterErr
				}
				if err == nil {
					value = durationNanos
				}
			}
		}
	default:
	}
	return value, nil
}

func validateValueType(name string, value any, fieldType enumspb.IndexedValueType) (any, error) {
	switch fieldType {
	case enumspb.INDEXED_VALUE_TYPE_INT, enumspb.INDEXED_VALUE_TYPE_DOUBLE:
		switch v := value.(type) {
		case int64, float64:
		// nothing to do
		case string:
			// ES can do implicit casting if the value is numeric
			if _, err := strconv.ParseFloat(v, 64); err != nil {
				return nil, query.NewConverterError(
					"invalid value for search attribute %s of type %s: %#v", name, fieldType.String(), value)
			}
		default:
			return nil, query.NewConverterError(
				"invalid value for search attribute %s of type %s: %#v", name, fieldType.String(), value)
		}
	case enumspb.INDEXED_VALUE_TYPE_BOOL:
		switch value.(type) {
		case bool:
		// nothing to do
		default:
			return nil, query.NewConverterError(
				"invalid value for search attribute %s of type %s: %#v", name, fieldType.String(), value)
		}
	case enumspb.INDEXED_VALUE_TYPE_DATETIME:
		switch v := value.(type) {
		case int64:
			value = time.Unix(0, v).UTC().Format(time.RFC3339Nano)
		case string:
			if _, err := time.Parse(time.RFC3339Nano, v); err != nil {
				return nil, query.NewConverterError(
					"invalid value for search attribute %s of type %s: %#v", name, fieldType.String(), value)
			}
		default:
			return nil, query.NewConverterError(
				"invalid value for search attribute %s of type %s: %#v", name, fieldType.String(), value)
		}
	}
	return value, nil
}
