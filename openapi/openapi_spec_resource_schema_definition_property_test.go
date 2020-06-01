package openapi

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTerraformCompliantPropertyName(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that has a name and not preferred name and name is compliant already", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "compliant_prop_name",
			Type: typeString,
		}
		Convey("When getTerraformCompliantPropertyName method is called", func() {
			compliantName := s.getTerraformCompliantPropertyName()
			Convey("Then the resulting name should be terraform compliant", func() {
				So(compliantName, ShouldEqual, "compliant_prop_name")
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that has a name and not preferred name and name is NOT compliant", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "nonCompliantName",
			Type: typeString,
		}
		Convey("When getTerraformCompliantPropertyName method is called", func() {
			compliantName := s.getTerraformCompliantPropertyName()
			Convey("Then the resulting name should be terraform compliant", func() {
				So(compliantName, ShouldEqual, "non_compliant_name")
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that has a name AND a preferred name and name is compliant", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:          "compliant_prop_name",
			PreferredName: "preferred_compliant_name",
			Type:          typeString,
		}
		Convey("When getTerraformCompliantPropertyName method is called", func() {
			compliantName := s.getTerraformCompliantPropertyName()
			Convey("Then the resulting name should be the preferred name", func() {
				So(compliantName, ShouldEqual, "preferred_compliant_name")
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that has a name AND a preferred name and preferred name is NOT compliant", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:          "compliant_prop_name",
			PreferredName: "preferredNonCompliantName",
			Type:          typeString,
		}
		Convey("When getTerraformCompliantPropertyName method is called", func() {
			compliantName := s.getTerraformCompliantPropertyName()
			Convey("Then the resulting name should be preferred name verbatim", func() {
				// If preferred name is set, whether the value is compliant or not that will be the value configured for
				// the terraform schema property. If the name is not terraform name compliant, Terraform will complain about
				// it at runtime
				So(compliantName, ShouldEqual, "preferredNonCompliantName")
			})
		})
	})
}

func TestIsPrimitiveProperty(t *testing.T) {
	testCases := []struct {
		name                 string
		specSchemaDefinition specSchemaDefinitionProperty
		expectedResult       bool
	}{
		{
			name: "pecSchemaDefinitionProperty that is a primitive string",
			specSchemaDefinition: specSchemaDefinitionProperty{
				Name: "primitive_property",
				Type: typeString,
			},
			expectedResult: true,
		},
		{
			name: "pecSchemaDefinitionProperty that is a primitive int",
			specSchemaDefinition: specSchemaDefinitionProperty{
				Name: "primitive_property",
				Type: typeInt,
			},
			expectedResult: true,
		},
		{
			name: "pecSchemaDefinitionProperty that is a primitive float",
			specSchemaDefinition: specSchemaDefinitionProperty{
				Name: "primitive_property",
				Type: typeFloat,
			},
			expectedResult: true,
		},
		{
			name: "pecSchemaDefinitionProperty that is a primitive bool",
			specSchemaDefinition: specSchemaDefinitionProperty{
				Name: "primitive_property",
				Type: typeBool,
			},
			expectedResult: true,
		},
		{
			name: "pecSchemaDefinitionProperty that is not a primitive",
			specSchemaDefinition: specSchemaDefinitionProperty{
				Name: "primitive_property",
				Type: typeObject,
			},
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		isPrimitiveProperty := tc.specSchemaDefinition.isPrimitiveProperty()
		assert.Equal(t, tc.expectedResult, isPrimitiveProperty, tc.name)
	}
}

func TestIsPropertyNamedID(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is PropertyNamedID", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: idDefaultPropertyName,
			Type: typeString,
		}
		Convey("When isPropertyNamedID method is called", func() {
			isPropertyNamedStatus := s.isPropertyNamedID()
			Convey("Then the resulted bool should be true", func() {
				So(isPropertyNamedStatus, ShouldBeTrue)
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that is PropertyNamedID with no compliant name", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "ID",
			Type: typeString,
		}
		Convey("When isPropertyNamedID method is called", func() {
			isPropertyNamedStatus := s.isPropertyNamedID()
			Convey("Then the resulted bool should be true", func() {
				So(isPropertyNamedStatus, ShouldBeTrue)
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that is NOT PropertyNamedID", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "some_other_property_name",
			Type:     typeString,
			Required: false,
		}
		Convey("When isPropertyNamedID method is called", func() {
			isPropertyNamedStatus := s.isPropertyNamedID()
			Convey("Then the resulted bool should be false", func() {
				So(isPropertyNamedStatus, ShouldBeFalse)
			})
		})
	})
}

func TestIsPropertyNamedStatus(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is PropertyNamedStatus", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: statusDefaultPropertyName,
			Type: typeString,
		}
		Convey("When isPropertyNamedStatus method is called", func() {
			isPropertyNamedStatus := s.isPropertyNamedStatus()
			Convey("Then the resulted bool should be true", func() {
				So(isPropertyNamedStatus, ShouldBeTrue)
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that is PropertyNamedStatus with no compliant name", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "Status",
			Type: typeString,
		}
		Convey("When isPropertyNamedStatus method is called", func() {
			isPropertyNamedStatus := s.isPropertyNamedStatus()
			Convey("Then the resulted bool should be true", func() {
				So(isPropertyNamedStatus, ShouldBeTrue)
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that is NOT PropertyNamedStatus", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "some_other_property_name",
			Type:     typeString,
			Required: false,
		}
		Convey("When isPropertyNamedStatus method is called", func() {
			isPropertyNamedStatus := s.isPropertyNamedStatus()
			Convey("Then the resulted bool should be false", func() {
				So(isPropertyNamedStatus, ShouldBeFalse)
			})
		})
	})
}

func TestIsObjectProperty(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is ObjectProperty", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "object_prop",
			Type: typeObject,
		}
		Convey("When isObjectProperty method is called", func() {
			isArrayProperty := s.isObjectProperty()
			Convey("Then the resulted bool should be true", func() {
				So(isArrayProperty, ShouldBeTrue)
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that is NOT ObjectProperty", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: false,
		}
		Convey("When isObjectProperty method is called", func() {
			isArrayProperty := s.isObjectProperty()
			Convey("Then the resulted bool should be false", func() {
				So(isArrayProperty, ShouldBeFalse)
			})
		})
	})
}

func TestIsArrayProperty(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is ArrayProperty", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeList,
			Required: true,
		}
		Convey("When isArrayTypeProperty method is called", func() {
			isArrayProperty := s.isArrayProperty()
			Convey("Then the resulted bool should be true", func() {
				So(isArrayProperty, ShouldBeTrue)
			})
		})
	})

	Convey("Given a specSchemaDefinitionProperty that is NOT ArrayProperty", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: false,
		}
		Convey("When isArrayTypeProperty method is called", func() {
			isArrayProperty := s.isArrayProperty()
			Convey("Then the resulted bool should be false", func() {
				So(isArrayProperty, ShouldBeFalse)
			})
		})
	})
}

func TestIsReadOnly(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is readOnly", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			ReadOnly: true,
		}
		Convey("When isReadOnly method is called", func() {
			isOptional := s.isReadOnly()
			Convey("Then the resulted bool should be true", func() {
				So(isOptional, ShouldBeTrue)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that is NOT readOnly", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			ReadOnly: false,
		}
		Convey("When isReadOnly method is called", func() {
			isOptional := s.isReadOnly()
			Convey("Then the resulted bool should be false", func() {
				So(isOptional, ShouldBeFalse)
			})
		})
	})
}

func TestIsOptional(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is NOT Required", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: false,
		}
		Convey("When isOptional method is called", func() {
			isOptional := s.isOptional()
			Convey("Then the resulted bool should be true", func() {
				So(isOptional, ShouldBeTrue)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that is NOT Required", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: true,
		}
		Convey("When isOptional method is called", func() {
			isOptional := s.isOptional()
			Convey("Then the resulted bool should be false", func() {
				So(isOptional, ShouldBeFalse)
			})
		})
	})
}

func TestIsRequired(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is Required", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: true,
		}
		Convey("When isRequired method is called", func() {
			isRequired := s.isRequired()
			Convey("Then the resulted bool should be true", func() {
				So(isRequired, ShouldBeTrue)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that is NOT Required", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: false,
		}
		Convey("When isRequired method is called", func() {
			isRequired := s.isRequired()
			Convey("Then the resulted bool should be false", func() {
				So(isRequired, ShouldBeFalse)
			})
		})
	})
}

func TestShouldIgnoreArrayItemsOrder(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is a typeList and where the 'x-terraform-ignore-order' ext is set to true", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:             "array_prop",
			Type:             typeList,
			IgnoreItemsOrder: true,
		}
		Convey("When shouldIgnoreArrayItemsOrder method is called", func() {
			isRequired := s.shouldIgnoreArrayItemsOrder()
			Convey("Then the resulted bool should be true", func() {
				So(isRequired, ShouldBeTrue)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that is a typeList and where the 'x-terraform-ignore-order' ext is set to false", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:             "array_prop",
			Type:             typeList,
			IgnoreItemsOrder: false,
		}
		Convey("When shouldIgnoreArrayItemsOrder method is called", func() {
			isRequired := s.shouldIgnoreArrayItemsOrder()
			Convey("Then the resulted bool should be false", func() {
				So(isRequired, ShouldBeFalse)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that is a typeList and where the 'x-terraform-ignore-order' ext is NOT set", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "array_prop",
			Type: typeList,
		}
		Convey("When shouldIgnoreArrayItemsOrder method is called", func() {
			isRequired := s.shouldIgnoreArrayItemsOrder()
			Convey("Then the resulted bool should be false", func() {
				So(isRequired, ShouldBeFalse)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that is NOT a typeList where the 'x-terraform-ignore-order' ext is set", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:             "string_prop",
			Type:             typeString,
			IgnoreItemsOrder: true,
		}
		Convey("When shouldIgnoreArrayItemsOrder method is called", func() {
			isRequired := s.shouldIgnoreArrayItemsOrder()
			Convey("Then the resulted bool should be false", func() {
				So(isRequired, ShouldBeFalse)
			})
		})
	})
}

func TestSchemaDefinitionPropertyIsComputed(t *testing.T) {
	Convey("Given a specSchemaDefinitionProperty that is optional and readonly", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: false,
			ReadOnly: true,
		}
		Convey("When isComputed method is called", func() {
			isReadOnly := s.isComputed()
			Convey("Then the resulted bool should be true", func() {
				So(isReadOnly, ShouldBeTrue)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that is optional, NOT readonly BUT is optional-computed", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: false,
			ReadOnly: false,
			Computed: true,
			Default:  nil,
		}
		Convey("When isComputed method is called", func() {
			isReadOnly := s.isComputed()
			Convey("Then the resulted bool should be true", func() {
				So(isReadOnly, ShouldBeTrue)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that NOT optional", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			Required: true,
		}
		Convey("When isComputed method is called", func() {
			isReadOnly := s.isComputed()
			Convey("Then the resulted bool should be false", func() {
				So(isReadOnly, ShouldBeFalse)
			})
		})
	})
	Convey("Given a specSchemaDefinitionProperty that is NOT readonly", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			ReadOnly: false,
		}
		Convey("When isComputed method is called", func() {
			isReadOnly := s.isComputed()
			Convey("Then the resulted bool should be false", func() {
				So(isReadOnly, ShouldBeFalse)
			})
		})
	})
}

func TestSchemaDefinitionPropertyIsOptionalComputed(t *testing.T) {
	Convey("Given a property that is optional, not readOnly, is computed and does not have a default value (optional-computed of property where value is not known at plan time)", t, func() {
		s := &specSchemaDefinitionProperty{
			Type:     typeString,
			Required: false,
			ReadOnly: false,
			Computed: true,
			Default:  nil,
		}
		Convey("When isOptionalComputed method is called", func() {
			isOptionalComputed := s.isOptionalComputed()
			Convey("Then value returned should be true", func() {
				So(isOptionalComputed, ShouldBeTrue)
			})
		})
	})
	Convey("Given a property that is not optional", t, func() {
		s := &specSchemaDefinitionProperty{
			Type:     typeString,
			Required: true,
		}
		Convey("When isOptionalComputed method is called", func() {
			isOptionalComputed := s.isOptionalComputed()
			Convey("Then value returned should be false", func() {
				So(isOptionalComputed, ShouldBeFalse)
			})
		})
	})
	Convey("Given a property that is optional but readOnly", t, func() {
		s := &specSchemaDefinitionProperty{
			Type:     typeString,
			Required: false,
			ReadOnly: true,
		}
		Convey("When isOptionalComputed method is called", func() {
			isOptionalComputed := s.isOptionalComputed()
			Convey("Then value returned should be false", func() {
				So(isOptionalComputed, ShouldBeFalse)
			})
		})
	})
	Convey("Given a property that is optional, not readOnly and it's not computed (purely optional use case)", t, func() {
		s := &specSchemaDefinitionProperty{
			Type:     typeString,
			Required: false,
			ReadOnly: false,
			Computed: false,
			Default:  nil,
		}
		Convey("When isOptionalComputed method is called", func() {
			isOptionalComputed := s.isOptionalComputed()
			Convey("Then value returned should be false", func() {
				So(isOptionalComputed, ShouldBeFalse)
			})
		})
	})
	Convey("Given a property that is optional, not readOnly, computed but has a default value (optional-computed use case, but as far as terraform is concerned the default will be set om the terraform schema, making it available at plan time - this is by design in terraform)", t, func() {
		s := &specSchemaDefinitionProperty{
			Type:     typeString,
			Required: false,
			ReadOnly: false,
			Computed: true,
			Default:  "something",
		}
		Convey("When isOptionalComputed method is called", func() {
			isOptionalComputed := s.isOptionalComputed()
			Convey("Then value returned should be false", func() {
				So(isOptionalComputed, ShouldBeFalse)
			})
		})
	})
}

func TestTerraformType(t *testing.T) {
	Convey("Given a swagger schema definition that has a property of type string", t, func() {
		s := &specSchemaDefinitionProperty{
			Type: typeString,
		}
		Convey("When terraformType method is called", func() {
			valueType, err := s.terraformType()
			Convey("Then error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And value type should be string", func() {
				So(valueType, ShouldEqual, schema.TypeString)
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type int", t, func() {
		s := &specSchemaDefinitionProperty{
			Type: typeInt,
		}
		Convey("When terraformType method is called", func() {
			valueType, err := s.terraformType()
			Convey("Then error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And valye type should be int", func() {
				So(valueType, ShouldEqual, schema.TypeInt)
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type float", t, func() {
		s := &specSchemaDefinitionProperty{
			Type: typeFloat,
		}
		Convey("When terraformType method is called", func() {
			valueType, err := s.terraformType()
			Convey("Then error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And valye type should be float", func() {
				So(valueType, ShouldEqual, schema.TypeFloat)
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type bool", t, func() {
		s := &specSchemaDefinitionProperty{
			Type: typeBool,
		}
		Convey("When terraformType method is called", func() {
			valueType, err := s.terraformType()
			Convey("Then error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And valye type should be bool", func() {
				So(valueType, ShouldEqual, schema.TypeBool)
			})
		})
	})
	Convey("Given a swagger schema definition that has an unsupported property type", t, func() {
		s := &specSchemaDefinitionProperty{
			Type: "unsupported",
		}
		Convey("When terraformType method is called", func() {
			valueType, err := s.terraformType()
			Convey("Then error returned should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("And value type should be invalid", func() {
				So(valueType, ShouldEqual, schema.TypeInvalid)
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type object", t, func() {
		s := &specSchemaDefinitionProperty{
			Type: typeObject,
		}
		Convey("When terraformType method is called", func() {
			valueType, err := s.terraformType()
			Convey("Then error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And valye type should be map", func() {
				So(valueType, ShouldEqual, schema.TypeMap)
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type list", t, func() {
		s := &specSchemaDefinitionProperty{
			Type: typeList,
		}
		Convey("When terraformType method is called", func() {
			valueType, err := s.terraformType()
			Convey("Then error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And valye type should be int", func() {
				So(valueType, ShouldEqual, schema.TypeList)
			})
		})
	})
}

func TestIsTerraformListOfSimpleValues(t *testing.T) {
	Convey("Given a swagger schema definition that has a property of type 'list' with elements of type string", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "list_prop",
			Type:           typeList,
			ArrayItemsType: typeString,
		}
		Convey("When isTerraformListOfSimpleValues method is called", func() {
			isTerraformListOfSimpleValues, listSchema := s.isTerraformListOfSimpleValues()
			Convey("Then the result should be true", func() {
				So(isTerraformListOfSimpleValues, ShouldBeTrue)
			})
			Convey("And the returned schema should be of tupe schema.Schema", func() {
				So(reflect.TypeOf(*listSchema), ShouldEqual, reflect.TypeOf(schema.Schema{}))
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type 'list' with elements of type int", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "list_prop",
			Type:           typeList,
			ArrayItemsType: typeInt,
		}
		Convey("When isTerraformListOfSimpleValues method is called", func() {
			isTerraformListOfSimpleValues, listSchema := s.isTerraformListOfSimpleValues()
			Convey("Then the result should be true", func() {
				So(isTerraformListOfSimpleValues, ShouldBeTrue)
			})
			Convey("And the returned schema should be of tupe schema.Schema", func() {
				So(reflect.TypeOf(*listSchema), ShouldEqual, reflect.TypeOf(schema.Schema{}))
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type 'list' with elements of type float", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "list_prop",
			Type:           typeList,
			ArrayItemsType: typeFloat,
		}
		Convey("When isTerraformListOfSimpleValues method is called", func() {
			isTerraformListOfSimpleValues, listSchema := s.isTerraformListOfSimpleValues()
			Convey("Then the result should be true", func() {
				So(isTerraformListOfSimpleValues, ShouldBeTrue)
			})
			Convey("And the returned schema should be of tupe schema.Schema", func() {
				So(reflect.TypeOf(*listSchema), ShouldEqual, reflect.TypeOf(schema.Schema{}))
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type 'list' with elements of type bool", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "list_prop",
			Type:           typeList,
			ArrayItemsType: typeBool,
		}
		Convey("When isTerraformListOfSimpleValues method is called", func() {
			isTerraformListOfSimpleValues, listSchema := s.isTerraformListOfSimpleValues()
			Convey("Then the result should be true", func() {
				So(isTerraformListOfSimpleValues, ShouldBeTrue)
			})
			Convey("And the returned schema should be of tupe schema.Schema", func() {
				So(reflect.TypeOf(*listSchema), ShouldEqual, reflect.TypeOf(schema.Schema{}))
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type 'list' with non primitive element", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "list_prop",
			Type:           typeList,
			ArrayItemsType: typeObject,
		}
		Convey("When isTerraformListOfSimpleValues method is called", func() {
			isTerraformListOfSimpleValues, listSchema := s.isTerraformListOfSimpleValues()
			Convey("Then the result should be true", func() {
				So(isTerraformListOfSimpleValues, ShouldBeFalse)
			})
			Convey("And the returned schema should be of tupe schema.Schema", func() {
				So(listSchema, ShouldBeNil)
			})
		})
	})
}

func TestTerraformObjectSchema(t *testing.T) {
	Convey("Given a swagger schema definition that has a property of type 'object'", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "object_prop",
			Type: typeObject,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Name: "protocol",
						Type: typeString,
					},
				},
			},
		}
		Convey("When terraformObjectSchema method is called", func() {
			tfPropSchema, err := s.terraformObjectSchema()
			Convey("Then the resulted tfPropSchema should be of type string too", func() {
				So(err, ShouldBeNil)
				So(reflect.TypeOf(*tfPropSchema), ShouldEqual, reflect.TypeOf(schema.Resource{}))
				So(tfPropSchema.Schema, ShouldContainKey, "protocol")
			})
		})
	})
	Convey("Given a swagger schema definition that has a property of type 'list' and arrays items type object", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "array_prop",
			Type:           typeList,
			ArrayItemsType: typeObject,
			ReadOnly:       false,
			Required:       true,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Name: "protocol",
						Type: typeString,
					},
				},
			},
		}
		Convey("When terraformObjectSchema method is called", func() {
			tfPropSchema, err := s.terraformObjectSchema()
			Convey("Then the resulted tfPropSchema should be of type string too", func() {
				So(err, ShouldBeNil)
				So(reflect.TypeOf(*tfPropSchema), ShouldEqual, reflect.TypeOf(schema.Resource{}))
				So(tfPropSchema.Schema, ShouldContainKey, "protocol")
			})
		})
	})

	Convey("Given a swagger schema definition that has a non supported property type for building object schmea", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "prop",
			Type: typeString,
		}
		Convey("When terraformObjectSchema method is called", func() {
			_, err := s.terraformObjectSchema()
			Convey("Then the error returned should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("Then the error message returned should match the expected one", func() {
				So(err.Error(), ShouldEqual, "object schema can only be formed for types object or types list with elems of type object: found type='string' elemType='' instead")
			})
		})
	})

	Convey("Given a swagger schema definition that has a object property type for building object schema", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "prop",
			Type: typeObject,
		}
		Convey("When terraformObjectSchema method is called", func() {
			_, err := s.terraformObjectSchema()
			Convey("Then the error returned should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("Then the error message returned should match the expected one", func() {
				So(err.Error(), ShouldEqual, "missing spec schema definition for property 'prop' of type 'object'")
			})
		})
	})
}

func TestIsLegacyComplexObjectExtensionEnabled(t *testing.T) {
	testCases := []struct {
		name                              string
		inputSpecSchemaDefinitionProperty specSchemaDefinitionProperty
		expectedResult                    bool
	}{
		{
			name: "property that is an object and has the EnableLegacyComplexObjectBlockConfiguration enabled",
			inputSpecSchemaDefinitionProperty: specSchemaDefinitionProperty{
				Type: typeObject,
				EnableLegacyComplexObjectBlockConfiguration: true,
			},
			expectedResult: true,
		},
		{
			name: "property that is an object and has the EnableLegacyComplexObjectBlockConfiguration NOT enabled",
			inputSpecSchemaDefinitionProperty: specSchemaDefinitionProperty{
				Type: typeObject,
				EnableLegacyComplexObjectBlockConfiguration: false,
			},
			expectedResult: false,
		},
		{
			name: "property that is NOT an object and has the EnableLegacyComplexObjectBlockConfiguration NOT enabled",
			inputSpecSchemaDefinitionProperty: specSchemaDefinitionProperty{
				Type: typeString,
				EnableLegacyComplexObjectBlockConfiguration: false,
			},
			expectedResult: false,
		},
		{
			name: "property that is NOT an object but somehow has the EnableLegacyComplexObjectBlockConfiguration turned on",
			inputSpecSchemaDefinitionProperty: specSchemaDefinitionProperty{
				Type: typeString,
				EnableLegacyComplexObjectBlockConfiguration: true,
			},
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		shouldEnableLegacyComplexObjectBlockConfiguration := tc.inputSpecSchemaDefinitionProperty.isLegacyComplexObjectExtensionEnabled()
		assert.Equal(t, tc.expectedResult, shouldEnableLegacyComplexObjectBlockConfiguration, tc.name)
	}
}

func TestSpecSchemaDefinitionIsPropertyWithNestedObjects(t *testing.T) {
	testcases := []struct {
		name                         string
		schemaDefinitionPropertyType schemaDefinitionPropertyType
		specSchemaDefinition         *specSchemaDefinition
		expected                     bool
	}{
		{name: "swagger schema definition property that is not of type 'object'",
			schemaDefinitionPropertyType: typeBool,
			specSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type: typeString,
					},
				},
			},
			expected: false},
		{name: "swagger schema definition property that has nested objects",
			schemaDefinitionPropertyType: typeObject,
			specSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type: typeObject,
						SpecSchemaDefinition: &specSchemaDefinition{
							Properties: specSchemaDefinitionProperties{
								&specSchemaDefinitionProperty{
									Type: typeString,
								},
							},
						},
					},
				},
			},
			expected: true},
		{name: "swagger schema definition property that DOES NOT have nested objects",
			specSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type: typeString,
					},
				},
			},
			expected: false},
		{name: "spec definition property of type object that does not have a corresponding spec schema definition",
			schemaDefinitionPropertyType: typeObject,
			specSchemaDefinition:         nil,
			expected:                     false},
	}
	for _, tc := range testcases {
		s := &specSchemaDefinitionProperty{
			Type:                 tc.schemaDefinitionPropertyType,
			SpecSchemaDefinition: tc.specSchemaDefinition,
		}
		isPropertyWithNestedObjects := s.isPropertyWithNestedObjects()
		assert.Equal(t, tc.expected, isPropertyWithNestedObjects, tc.name)

	}

}

func TestTerraformSchema(t *testing.T) {
	Convey("Given a swagger schema definition that has two nested properties - one being a simple object and the other one a primitive", t, func() {
		expectedNestedObjectPropertyName := "nested_object1"
		s := &specSchemaDefinitionProperty{
			Name: "top_level_object",
			Type: typeObject,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type: typeObject,
						Name: expectedNestedObjectPropertyName,
						SpecSchemaDefinition: &specSchemaDefinition{
							Properties: specSchemaDefinitionProperties{
								&specSchemaDefinitionProperty{
									Type: typeString,
									Name: "string_property_1",
								},
							},
						},
					},
					&specSchemaDefinitionProperty{
						Type:          typeFloat,
						Name:          "nested_float2",
						PreferredName: "nested_float_2",
					},
				},
			}}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulting tfPropSchema should have a top level that is a 1 element list (workaround for object property with nested object)", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
				So(tfPropSchema.MaxItems, ShouldEqual, 1)
			})
			Convey("And the returned terraform schema contains the 'nested_object1' with the right configuration", func() {
				nestedObject1 := tfPropSchema.Elem.(*schema.Resource).Schema["nested_object1"]
				So(nestedObject1, ShouldNotBeNil)
				So(nestedObject1.Type, ShouldEqual, schema.TypeMap)
				So(nestedObject1.Elem.(*schema.Resource).Schema["string_property_1"].Type, ShouldEqual, schema.TypeString)
			})
			Convey("And the returned terraform schema contains the 'nested_float_2' with the right configuration", func() {
				nestedObject2 := tfPropSchema.Elem.(*schema.Resource).Schema["nested_float_2"]
				So(nestedObject2.Type, ShouldEqual, schema.TypeFloat)
			})
		})
	})

	Convey("Given a swagger schema definition that has two nested simple object properties", t, func() {
		expectedNestedObjectPropertyName1 := "nested_object1"
		expectedNestedObjectPropertyName2 := "nested_object2"
		s := &specSchemaDefinitionProperty{
			Name: "top_level_object",
			Type: typeObject,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type: typeObject,
						Name: expectedNestedObjectPropertyName1,
						SpecSchemaDefinition: &specSchemaDefinition{
							Properties: specSchemaDefinitionProperties{
								&specSchemaDefinitionProperty{
									Type: typeString,
									Name: "string_property_1",
								},
							},
						},
					},
					&specSchemaDefinitionProperty{
						Type: typeObject,
						Name: expectedNestedObjectPropertyName2,
						SpecSchemaDefinition: &specSchemaDefinition{
							Properties: specSchemaDefinitionProperties{
								&specSchemaDefinitionProperty{
									Type: typeString,
									Name: "string_property_2",
								},
							},
						},
					},
				},
			}}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulting tfPropSchema should have a top level that is a 1 element list (workaround for object property with nested object)", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
				So(tfPropSchema.MaxItems, ShouldEqual, 1)
			})
			Convey("And the returned terraform schema contains the schema for the first nested object property with the right configuration", func() {
				nestedObject1 := tfPropSchema.Elem.(*schema.Resource).Schema[expectedNestedObjectPropertyName1]
				So(nestedObject1, ShouldNotBeNil)
				So(nestedObject1.Type, ShouldEqual, schema.TypeMap)
				So(nestedObject1.Elem.(*schema.Resource).Schema["string_property_1"].Type, ShouldEqual, schema.TypeString)
			})
			Convey("And the returned terraform schema contains the schema for the Second nested object property with the right configuration", func() {
				nestedObject2 := tfPropSchema.Elem.(*schema.Resource).Schema[expectedNestedObjectPropertyName2]
				So(nestedObject2, ShouldNotBeNil)
				So(nestedObject2.Type, ShouldEqual, schema.TypeMap)
				So(nestedObject2.Elem.(*schema.Resource).Schema["string_property_2"].Type, ShouldEqual, schema.TypeString)
			})
		})
	})

	Convey("Given a swagger schema definition of type object and a complex object nested into it", t, func() {
		complexObjectName := "complex_object_which_is_nested"
		s := &specSchemaDefinitionProperty{
			Name: "top_level_object",
			Type: typeObject,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type: typeObject,
						Name: complexObjectName,
						EnableLegacyComplexObjectBlockConfiguration: true,
						SpecSchemaDefinition: &specSchemaDefinition{
							Properties: specSchemaDefinitionProperties{
								&specSchemaDefinitionProperty{
									Type: typeString,
									Name: "string_property",
								},
								&specSchemaDefinitionProperty{
									Type:     typeInt,
									Name:     "int_property",
									ReadOnly: true,
								},
							},
						},
					},
				},
			}}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulting tfPropSchema should have a top level that is a 1 element list (workaround for object property with nested object)", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
				So(tfPropSchema.MaxItems, ShouldEqual, 1)
			})
			Convey("And the returned terraform schema contains the schema for the first nested AND complex object property with the right configuration --> typeList", func() {
				nestedAndComplexObj := tfPropSchema.Elem.(*schema.Resource).Schema[complexObjectName]
				So(nestedAndComplexObj, ShouldNotBeNil)
				So(nestedAndComplexObj.Type, ShouldEqual, schema.TypeList)
				So(nestedAndComplexObj.MaxItems, ShouldEqual, 1)
				So(nestedAndComplexObj.Elem.(*schema.Resource).Schema["string_property"].Type, ShouldEqual, schema.TypeString)
				So(nestedAndComplexObj.Elem.(*schema.Resource).Schema["int_property"].Type, ShouldEqual, schema.TypeInt)
				So(nestedAndComplexObj.Elem.(*schema.Resource).Schema["int_property"].Computed, ShouldBeTrue)
			})
		})
	})

	Convey("Given a swagger schema definition that contains a nested object", t, func() {
		complexObjectName := "complex_object_which_is_nested"
		s := &specSchemaDefinitionProperty{
			Name: "complex object",
			Type: typeObject,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type:                 typeObject,
						Name:                 complexObjectName,
						ReadOnly:             true,
						SpecSchemaDefinition: &specSchemaDefinition{},
					},
				},
			}}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulting tfPropSchema should have mapped the complex object as a 1 element list BECAUSE even if the complex object doesn't have EnableLegacyComplexObjectBlockConfiguration=true being nested == being complex object", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
				So(tfPropSchema.MaxItems, ShouldEqual, 1)
				So(tfPropSchema.Elem.(*schema.Resource).Schema[complexObjectName].Computed, ShouldBeTrue)
			})
		})
	})

	Convey("Given a swagger schema definition that contains a object with no SpecSchemaDefinition", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "object",
			Type: typeObject,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type: typeObject,
						Name: "the object",
					},
				},
			}}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then return an error because of the lack of SpecSchemaDefinition", func() {
				So(err, ShouldNotBeNil)
				So(tfPropSchema, ShouldBeNil)
				So(err.Error(), ShouldEqual, `missing spec schema definition for property 'the object' of type 'object'`)

			})
		})
	})

	Convey("Given a swagger schema definition tha is a complex object (EnableLegacyComplexObjectBlockConfiguration set to true)", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "top level complex object",
			Type: typeObject,
			EnableLegacyComplexObjectBlockConfiguration: true,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type:                 typeInt,
						Name:                 "my_int_prop",
						ReadOnly:             true,
						SpecSchemaDefinition: &specSchemaDefinition{},
					},
				},
			}}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulting tfPropSchema should have mapped the object as a 1 elem list BECAUSE it has EnableLegacyComplexObjectBlockConfiguration = true", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
				So(tfPropSchema.MaxItems, ShouldEqual, 1)
				So(tfPropSchema.Elem.(*schema.Resource).Schema["my_int_prop"].Computed, ShouldBeTrue)
			})
		})
	})

	Convey("Given a swagger schema definition tha is a simple object (EnableLegacyComplexObjectBlockConfiguration not present or set to false)", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "top level complex object",
			Type: typeObject,
			//EnableLegacyComplexObjectBlockConfiguration: true, ==> This field is not present or set to false
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Type:                 typeInt,
						Name:                 "my_int_prop",
						ReadOnly:             true,
						SpecSchemaDefinition: &specSchemaDefinition{},
					},
				},
			}}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulting tfPropSchema should match the following using TypeMap as type", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeMap)
				So(tfPropSchema.Elem.(*schema.Resource).Schema["my_int_prop"].Computed, ShouldBeTrue)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'string' which is required", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "string_prop",
			Type:     typeString,
			ReadOnly: false,
			Required: true,
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted tfPropSchema should be of type string too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeString)
				So(tfPropSchema.Required, ShouldBeTrue)
			})
		})
	})
	Convey("Given a swagger schema definition that has an unsupported property type", t, func() {
		s := &specSchemaDefinitionProperty{
			Name: "rune_prop",
			Type: "unsupported",
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be nil", func() {
				So(err.Error(), ShouldEqual, "non supported type unsupported")
				So(tfPropSchema, ShouldBeNil)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'integer'", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "int_prop",
			Type:     typeInt,
			ReadOnly: false,
			Required: true,
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be of type int too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeInt)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'number'", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "number_prop",
			Type:     typeFloat,
			ReadOnly: false,
			Required: true,
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be of type float too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeFloat)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'boolean'", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "boolean_prop",
			Type:     typeBool,
			ReadOnly: false,
			Required: true,
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be of type int too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeBool)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'array' and the elems are type string", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "array_prop",
			Type:           typeList,
			ArrayItemsType: typeString,
			ReadOnly:       false,
			Required:       true,
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be of type array too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
			})
			Convey("And the array elements are of type string", func() {
				So(reflect.TypeOf(tfPropSchema.Elem).Elem(), ShouldEqual, reflect.TypeOf(schema.Schema{}))
				So(tfPropSchema.Elem.(*schema.Schema).Type, ShouldEqual, schema.TypeString)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'array' and the elems are type integer", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "array_prop",
			Type:           typeList,
			ArrayItemsType: typeInt,
			ReadOnly:       false,
			Required:       true,
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be of type array too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
			})
			Convey("And the array elements are of type int", func() {
				So(reflect.TypeOf(tfPropSchema.Elem).Elem(), ShouldEqual, reflect.TypeOf(schema.Schema{}))
				So(tfPropSchema.Elem.(*schema.Schema).Type, ShouldEqual, schema.TypeInt)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'array' and the elems are type number", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "array_prop",
			Type:           typeList,
			ArrayItemsType: typeFloat,
			ReadOnly:       false,
			Required:       true,
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be of type array too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
			})
			Convey("And the array elements are of type float", func() {
				So(reflect.TypeOf(tfPropSchema.Elem).Elem(), ShouldEqual, reflect.TypeOf(schema.Schema{}))
				So(tfPropSchema.Elem.(*schema.Schema).Type, ShouldEqual, schema.TypeFloat)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'array' and the elems are type bool", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "array_prop",
			Type:           typeList,
			ArrayItemsType: typeBool,
			ReadOnly:       false,
			Required:       true,
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be of type array too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
			})
			Convey("And the array elements are of type bool", func() {
				So(reflect.TypeOf(tfPropSchema.Elem).Elem(), ShouldEqual, reflect.TypeOf(schema.Schema{}))
				So(tfPropSchema.Elem.(*schema.Schema).Type, ShouldEqual, schema.TypeBool)
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'array' and the elems are type object", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:           "array_prop",
			Type:           typeList,
			ArrayItemsType: typeObject,
			ReadOnly:       false,
			Required:       true,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Name: "protocol",
						Type: typeString,
					},
				},
			},
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the resulted terraform property schema should be of type array too", func() {
				So(err, ShouldBeNil)
				So(tfPropSchema.Type, ShouldEqual, schema.TypeList)
			})
			Convey("And the array elements are of type object (resource object) containing the object schema properties", func() {
				So(reflect.TypeOf(tfPropSchema.Elem).Elem(), ShouldEqual, reflect.TypeOf(schema.Resource{}))
				So(tfPropSchema.Elem.(*schema.Resource).Schema, ShouldContainKey, "protocol")
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type 'array' and the elems are not set", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "array_prop",
			Type:     typeList,
			ReadOnly: false,
			Required: true,
		}
		Convey("When terraformSchema method is called", func() {
			_, err := s.terraformSchema()
			Convey("Then the error returned should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("Then the error message returned should be the expected one", func() {
				So(err.Error(), ShouldEqual, "object schema can only be formed for types object or types list with elems of type object: found type='list' elemType='' instead")
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type object and a ref pointing to the schema", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "object_prop",
			Type:     typeObject,
			ReadOnly: false,
			Required: true,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Name: "message",
						Type: typeString,
					},
				},
			},
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the tf resource schema returned should not be nil", func() {
				So(tfPropSchema, ShouldNotBeNil)
			})
			Convey("And the tf resource schema returned should contained the sub property - 'message'", func() {
				So(tfPropSchema.Elem.(*schema.Resource).Schema, ShouldContainKey, "message")
			})
		})
	})

	Convey("Given a swagger schema definition that has a property of type object that has nested schema and property named id", t, func() {
		s := &specSchemaDefinitionProperty{
			Name:     "object_prop",
			Type:     typeObject,
			ReadOnly: false,
			Required: true,
			SpecSchemaDefinition: &specSchemaDefinition{
				Properties: specSchemaDefinitionProperties{
					&specSchemaDefinitionProperty{
						Name: "id",
						Type: typeString,
					},
					&specSchemaDefinitionProperty{
						Name: "message",
						Type: typeString,
					},
				},
			},
		}
		Convey("When terraformSchema method is called", func() {
			tfPropSchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the tf resource schema returned should not be nil", func() {
				So(tfPropSchema, ShouldNotBeNil)
			})
			Convey("And the tf resource schema returned should contain the sub property - 'message'", func() {
				So(tfPropSchema.Elem.(*schema.Resource).Schema, ShouldContainKey, "message")
			})
			Convey("And the tf resource schema returned should contain the sub property - 'id' and should not be ignored", func() {
				So(tfPropSchema.Elem.(*schema.Resource).Schema, ShouldContainKey, "id")
			})
		})
	})

	Convey("Given a string schemaDefinitionProperty that is required, not readOnly, forceNew, sensitive, not immutable and has a default value", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", true, false, false, true, true, false, false, false, "default value")
		Convey("When terraformSchema is called with a schema definition property that is required, force new, sensitive and has a default value", func() {
			terraformPropertySchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the schema returned should be configured as required", func() {
				So(terraformPropertySchema.Required, ShouldBeTrue)
			})
			Convey("And the schema returned should not be configured as optional", func() {
				So(terraformPropertySchema.Optional, ShouldBeFalse)
			})
			Convey("And the schema returned should be configured as NOT computed", func() {
				So(terraformPropertySchema.Computed, ShouldBeFalse)
			})
			Convey("And the schema returned should be configured as force new", func() {
				So(terraformPropertySchema.ForceNew, ShouldBeTrue)
			})
			Convey("And the schema returned should be configured as sensitive", func() {
				So(terraformPropertySchema.Sensitive, ShouldBeTrue)
			})
			Convey("And the schema returned should have a default value (note: terraform will complain in this case at runtime since required properties cannot have default values)", func() {
				So(terraformPropertySchema.Default, ShouldEqual, s.Default)
			})
			Convey("And the schema returned should be configured with a validate function", func() {
				So(terraformPropertySchema.ValidateFunc, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a schemaDefinitionProperty that is readOnly and does not have a default value (meaning the value is not known at plan time)", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", false, true, true, false, false, false, false, false, "")
		Convey("When terraformSchema is called with a schema definition property that is readonly", func() {
			terraformPropertySchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the schema returned should be configured as not required", func() {
				So(terraformPropertySchema.Required, ShouldBeFalse)
			})
			Convey("And the schema returned should be configured as optional", func() {
				So(terraformPropertySchema.Optional, ShouldBeTrue)
			})
			Convey("And the schema returned should be configured as computed", func() {
				So(terraformPropertySchema.Computed, ShouldBeTrue)
			})
			Convey("And the schema returned should be configured with a validate function", func() {
				So(terraformPropertySchema.ValidateFunc, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a schemaDefinitionProperty that is readOnly and does have a default value (meaning the default value is known at plan time)", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", false, true, true, false, false, false, false, false, "some value")
		Convey("When terraformSchema is called with a schema definition property that is readonly", func() {
			terraformPropertySchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the schema returned should be configured as not required", func() {
				So(terraformPropertySchema.Required, ShouldBeFalse)
			})
			Convey("And the schema returned should be configured as optional", func() {
				So(terraformPropertySchema.Optional, ShouldBeTrue)
			})
			Convey("And the schema returned should be configured as computed", func() {
				So(terraformPropertySchema.Computed, ShouldBeTrue)
			})
			Convey("And the schema returned should not be configured since the property is readOnly", func() {
				So(terraformPropertySchema.Default, ShouldBeNil)
			})
			Convey("And the schema returned should be configured with a validate function", func() {
				So(terraformPropertySchema.ValidateFunc, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a schemaDefinitionProperty that is optional computed and does not have a default value (meaning the value is not known at plan time)", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", false, false, true, false, false, false, false, false, nil)
		Convey("When terraformSchema is called with a schema definition property that is optional computed", func() {
			terraformPropertySchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the schema returned should be configured as not required", func() {
				So(terraformPropertySchema.Required, ShouldBeFalse)
			})
			Convey("And the schema returned should be configured as optional", func() {
				So(terraformPropertySchema.Optional, ShouldBeTrue)
			})
			Convey("And the schema returned should be configured as computed", func() {
				So(terraformPropertySchema.Computed, ShouldBeTrue)
			})
			Convey("And the schema returned should not be configured with a default value", func() {
				So(terraformPropertySchema.Default, ShouldBeNil)
			})
			Convey("And the schema returned should be configured with a validate function", func() {
				So(terraformPropertySchema.ValidateFunc, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a schemaDefinitionProperty that is optional computed and does have a default value (meaning the value is known at plan time)", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", false, false, true, false, false, false, false, false, "some known value")
		Convey("When terraformSchema is called with a schema definition property that is optional computed", func() {
			terraformPropertySchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the schema returned should be configured as not required", func() {
				So(terraformPropertySchema.Required, ShouldBeFalse)
			})
			Convey("And the schema returned should be configured as optional", func() {
				So(terraformPropertySchema.Optional, ShouldBeTrue)
			})
			Convey("And the schema returned should not be configured as computed since in this case terraform will make use of the default value as input for the user. This makes the default value visible at plan time", func() {
				So(terraformPropertySchema.Computed, ShouldBeFalse)
			})
			Convey("And the schema returned should be configured with a default value", func() {
				So(terraformPropertySchema.Default, ShouldNotBeNil)
			})
			Convey("And the schema returned should be configured with a validate function", func() {
				So(terraformPropertySchema.ValidateFunc, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a schemaDefinitionProperty that is forceNew and immutable ", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", false, false, false, true, false, true, false, false, "")
		Convey("When terraformSchema is called with a schema definition property that validation fails due to immutable and forceNew set", func() {
			terraformPropertySchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the schema returned should be configured with a validate function", func() {
				So(terraformPropertySchema.ValidateFunc, ShouldNotBeNil)
			})
			Convey("And the schema validate function should return an error ", func() {
				_, err := terraformPropertySchema.ValidateFunc(nil, "")
				So(err, ShouldNotBeNil)
				So(err, ShouldNotBeEmpty)
				So(err[0].Error(), ShouldContainSubstring, "property 'propertyName' is configured as immutable and can not be configured with forceNew too")
			})
		})
	})

	Convey("Given a schemaDefinitionProperty that is readOnly and required", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", true, true, false, false, false, false, false, false, nil)
		Convey("When terraformSchema is called with a schema definition property that validation fails due to required and computed set", func() {
			terraformPropertySchema, err := s.terraformSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the schema returned should be configured as required", func() {
				So(terraformPropertySchema.Required, ShouldBeTrue)
			})
			Convey("And the schema returned should not be configured as computed as it's not optional property", func() {
				So(terraformPropertySchema.Computed, ShouldBeFalse)
			})
			Convey("And the schema returned should be configured with a validate function", func() {
				So(terraformPropertySchema.ValidateFunc, ShouldNotBeNil)
			})
			Convey("And the schema validate function should return an error ", func() {
				_, err := terraformPropertySchema.ValidateFunc(nil, "")
				So(err, ShouldNotBeNil)
				So(err, ShouldNotBeEmpty)
				So(err[0].Error(), ShouldContainSubstring, "property 'propertyName' is configured as required and can not be configured as computed too")
			})
		})
	})
}

func TestValidateFunc(t *testing.T) {

	Convey("Given a schemaDefinitionProperty that is computed and has a default value set", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", false, true, false, false, false, false, false, false, "defaultValue")
		Convey("When validateFunc is called with a schema definition property", func() {
			Convey("And the schema returned should be configured with a validate function", func() {
				So(s.validateFunc(), ShouldNotBeNil)
			})
			Convey("And the schema validate function should return successfully", func() {
				_, err := s.validateFunc()(nil, "")
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given a schemaDefinitionProperty that is forceNew and immutable ", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", false, false, false, true, false, true, false, false, "")
		Convey("When validateFunc is called with a schema definition property", func() {
			Convey("And the schema returned should be configured with a validate function", func() {
				So(s.validateFunc(), ShouldNotBeNil)
			})
			Convey("And the schema validate function should return an error due to immutable and forceNew set", func() {
				_, err := s.validateFunc()(nil, "")
				So(err, ShouldNotBeNil)
				So(err, ShouldNotBeEmpty)
				So(err[0].Error(), ShouldContainSubstring, "property 'propertyName' is configured as immutable and can not be configured with forceNew too")
			})
		})
	})

	Convey("Given a schemaDefinitionProperty that is computed and required", t, func() {
		s := newStringSchemaDefinitionProperty("propertyName", "", true, true, false, false, false, false, false, false, nil)
		Convey("When validateFunc is called with a schema definition property", func() {
			Convey("And the schema returned should be configured with a validate function", func() {
				So(s.validateFunc(), ShouldNotBeNil)
			})
			Convey("And the schema validate function should return an error due to required and computed set", func() {
				_, err := s.validateFunc()(nil, "")
				So(err, ShouldNotBeNil)
				So(err, ShouldNotBeEmpty)
				So(err[0].Error(), ShouldContainSubstring, "property 'propertyName' is configured as required and can not be configured as computed too")
			})
		})
	})
}

func TestEqualItems(t *testing.T) {
	testCases := []struct {
		name               string
		schemaDefProp      specSchemaDefinitionProperty
		propertyType       schemaDefinitionPropertyType
		arrayItemsPropType schemaDefinitionPropertyType
		inputItem          interface{}
		remoteItem         interface{}
		expectedOutput     bool
	}{
		// String use cases
		{
			name:           "string input value matches string remote value",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeString},
			inputItem:      "inputVal1",
			remoteItem:     "inputVal1",
			expectedOutput: true,
		},
		{
			name:           "string input value doesn't match string remote value",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeString},
			inputItem:      "inputVal1",
			remoteItem:     "inputVal2",
			expectedOutput: false,
		},
		// Integer use cases
		{
			name:           "int input value matches int remote value",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeInt},
			inputItem:      1,
			remoteItem:     1,
			expectedOutput: true,
		},
		{
			name:           "int input value doesn't match int remote value",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeInt},
			inputItem:      1,
			remoteItem:     2,
			expectedOutput: false,
		},
		// Float use cases
		{
			name:           "float input value matches float remote value",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeFloat},
			inputItem:      1.0,
			remoteItem:     1.0,
			expectedOutput: true,
		},
		{
			name:           "float input value doesn't match float remote value",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeFloat},
			inputItem:      1.0,
			remoteItem:     2.0,
			expectedOutput: false,
		},
		// Bool use cases
		{
			name:           "bool input value matches bool remote value",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeBool},
			inputItem:      true,
			remoteItem:     true,
			expectedOutput: true,
		},
		{
			name:           "bool input value doesn't match bool remote value",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeBool},
			inputItem:      true,
			remoteItem:     false,
			expectedOutput: false,
		},
		// List use cases
		{
			name: "list input value matches list remote value",
			schemaDefProp: specSchemaDefinitionProperty{
				Type:           typeList,
				ArrayItemsType: typeString,
			},
			inputItem:      []interface{}{"role1", "role2"},
			remoteItem:     []interface{}{"role1", "role2"},
			expectedOutput: true,
		},
		{
			name: "list input value doesn't match list remote value (same list length)",
			schemaDefProp: specSchemaDefinitionProperty{
				Type:           typeList,
				ArrayItemsType: typeString,
			},
			inputItem:      []interface{}{"role1", "role2"},
			remoteItem:     []interface{}{"role3", "role4"},
			expectedOutput: false,
		},
		{
			name: "list input value doesn't match list remote value (same list length band same items but unordered) but property is marked as ignore order",
			schemaDefProp: specSchemaDefinitionProperty{
				Type:             typeList,
				ArrayItemsType:   typeString,
				IgnoreItemsOrder: true,
			},
			inputItem:      []interface{}{"role1", "role2"},
			remoteItem:     []interface{}{"role2", "role1"},
			expectedOutput: true,
		},
		{
			name: "list input value doesn't match list remote value (same list length band) but property is marked as ignore order",
			schemaDefProp: specSchemaDefinitionProperty{
				Type:             typeList,
				ArrayItemsType:   typeString,
				IgnoreItemsOrder: true,
			},
			inputItem:      []interface{}{"role1", "role2"},
			remoteItem:     []interface{}{"role3", "role1"},
			expectedOutput: false,
		},
		{
			name: "list input value doesn't match list remote value (different list length)",
			schemaDefProp: specSchemaDefinitionProperty{
				Type:           typeList,
				ArrayItemsType: typeString,
			},
			inputItem:      []interface{}{"role1", "role2"},
			remoteItem:     []interface{}{"role1"},
			expectedOutput: false,
		},
		// Object use cases
		{
			name: "object input value matches object remote value",
			schemaDefProp: specSchemaDefinitionProperty{
				Type: typeObject,
				SpecSchemaDefinition: &specSchemaDefinition{
					Properties: specSchemaDefinitionProperties{
						&specSchemaDefinitionProperty{
							Name: "group",
							Type: typeString,
						},
					},
				},
			},
			inputItem:      map[string]interface{}{"group": "someGroup"},
			remoteItem:     map[string]interface{}{"group": "someGroup"},
			expectedOutput: true,
		},
		{
			name: "object input value doesn't match object remote value",
			schemaDefProp: specSchemaDefinitionProperty{
				Type: typeObject,
				SpecSchemaDefinition: &specSchemaDefinition{
					Properties: specSchemaDefinitionProperties{
						&specSchemaDefinitionProperty{
							Name: "group",
							Type: typeString,
						},
					},
				},
			},
			inputItem:      map[string]interface{}{"group": "someGroup"},
			remoteItem:     map[string]interface{}{"group": "someOtherGroup"},
			expectedOutput: false,
		},
		// Negative cases
		{
			name:           "string input value is not a string",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeString},
			inputItem:      1,
			remoteItem:     "inputVal1",
			expectedOutput: false,
		},
		{
			name:           "int input value is not an int",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeInt},
			inputItem:      "not_an_int",
			remoteItem:     1,
			expectedOutput: false,
		},
		{
			name:           "float input value is not a float",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeFloat},
			inputItem:      1.0,
			remoteItem:     "not_an_float",
			expectedOutput: false,
		},
		{
			name:           "bool input value is nto a bool",
			schemaDefProp:  specSchemaDefinitionProperty{Type: typeBool},
			inputItem:      true,
			remoteItem:     "not_a_bool",
			expectedOutput: false,
		},
		{
			name: "list input value is not a list",
			schemaDefProp: specSchemaDefinitionProperty{
				Type:           typeList,
				ArrayItemsType: typeString,
			},
			inputItem:      []interface{}{"role1", "role2"},
			remoteItem:     "not a list",
			expectedOutput: false,
		},
		{
			name: "object input value is not an object",
			schemaDefProp: specSchemaDefinitionProperty{
				Type: typeObject,
				SpecSchemaDefinition: &specSchemaDefinition{
					Properties: specSchemaDefinitionProperties{
						&specSchemaDefinitionProperty{
							Name: "group",
							Type: typeString,
						},
					},
				},
			},
			inputItem:      "not_an_object",
			remoteItem:     map[string]interface{}{"group": "someGroup"},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		output := tc.schemaDefProp.equal(tc.inputItem, tc.remoteItem)
		assert.Equal(t, tc.expectedOutput, output, tc.name)
	}
}

func TestValidateValueType(t *testing.T) {
	testCases := []struct {
		name           string
		item           interface{}
		itemKind       reflect.Kind
		expectedOutput bool
	}{
		// String use cases
		{
			name:           "expect string kind and item is a string",
			item:           "inputVal1",
			itemKind:       reflect.String,
			expectedOutput: true,
		},
		{
			name:           "expect string kind and item is NOT a string",
			item:           1,
			itemKind:       reflect.String,
			expectedOutput: false,
		},
		// Int use cases
		{
			name:           "expect int kind and item is a int",
			item:           1,
			itemKind:       reflect.Int,
			expectedOutput: true,
		},
		{
			name:           "expect int kind and item is NOT a int",
			item:           "not an int",
			itemKind:       reflect.Int,
			expectedOutput: false,
		},
		// Float use cases
		{
			name:           "expect float kind and item is a float",
			item:           1.0,
			itemKind:       reflect.Float64,
			expectedOutput: true,
		},
		{
			name:           "expect float kind and item is NOT a float",
			item:           "not a float",
			itemKind:       reflect.Float64,
			expectedOutput: false,
		},
		// Bool use cases
		{
			name:           "expect bool kind and item is a bool",
			item:           true,
			itemKind:       reflect.Bool,
			expectedOutput: true,
		},
		{
			name:           "expect bool kind and item is NOT a bool",
			item:           "not a bool",
			itemKind:       reflect.Bool,
			expectedOutput: false,
		},
		//  List use cases
		{
			name:           "expect slice kind and item is a slice",
			item:           []interface{}{"item1", "item2"},
			itemKind:       reflect.Slice,
			expectedOutput: true,
		},
		{
			name:           "expect slice kind and item is NOT a slice",
			item:           "not a slice",
			itemKind:       reflect.Slice,
			expectedOutput: false,
		},
		//  Object use cases
		{
			name:           "expect map kind and item is a map",
			item:           map[string]interface{}{"group": "someGroup"},
			itemKind:       reflect.Map,
			expectedOutput: true,
		},
		{
			name:           "expect map kind and item is NOT a map",
			item:           "not a map",
			itemKind:       reflect.Map,
			expectedOutput: false,
		},
	}
	for _, tc := range testCases {
		s := specSchemaDefinitionProperty{}
		output := s.validateValueType(tc.item, tc.itemKind)
		assert.Equal(t, tc.expectedOutput, output, tc.name)
	}
}

func Test_shouldIgnoreOrder(t *testing.T) {
	Convey("Given a scjema definition property that is a list and configured with ignore order", t, func() {
		p := &specSchemaDefinitionProperty{
			Type:             typeList,
			IgnoreItemsOrder: true,
		}
		Convey("When shouldIgnoreOrder is called", func() {
			shouldIgnoreOrder := p.shouldIgnoreOrder()
			Convey("Then the err returned should be true", func() {
				So(shouldIgnoreOrder, ShouldBeTrue)
			})
		})
	})
	Convey("Given a scjema definition property that is NOT a list", t, func() {
		p := &specSchemaDefinitionProperty{
			Type: typeString,
		}
		Convey("When shouldIgnoreOrder is called", func() {
			shouldIgnoreOrder := p.shouldIgnoreOrder()
			Convey("Then the err returned should be false", func() {
				So(shouldIgnoreOrder, ShouldBeFalse)
			})
		})
	})
	Convey("Given a scjema definition property that is a list but the ignore order is set to false", t, func() {
		p := &specSchemaDefinitionProperty{
			Type:             typeList,
			IgnoreItemsOrder: false,
		}
		Convey("When shouldIgnoreOrder is called", func() {
			shouldIgnoreOrder := p.shouldIgnoreOrder()
			Convey("Then the err returned should be false", func() {
				So(shouldIgnoreOrder, ShouldBeFalse)
			})
		})
	})
}

func Test_shouldUseLegacyTerraformSDKBlockApproachForComplexObjects(t *testing.T) {

	Convey("shouldUseLegacyTerraformSDKBlockApproachForComplexObject", t, func() {

		Convey("returns false", func() {
			Convey("Given a blank specSchemaDefinitionProperty "+
				"When shouldUseLegacyTerraformSDKBlockApproachForComplexObjects is called"+
				"Then it should return false with no error", func() {
				p := &specSchemaDefinitionProperty{}
				b := p.shouldUseLegacyTerraformSDKBlockApproachForComplexObjects()
				So(b, ShouldBeFalse)
			})

			Convey("Given a specSchemaDefinitionProperty with Type of 'boolean' and nothing else"+
				"When shouldUseLegacyTerraformSDKBlockApproachForComplexObjects is called"+
				"Then it should return false with no error", func() {
				p := &specSchemaDefinitionProperty{Type: typeBool}
				b := p.shouldUseLegacyTerraformSDKBlockApproachForComplexObjects()
				So(b, ShouldBeFalse)
			})

			Convey("Given a specSchemaDefinitionProperty with Type of 'object' and a blank SpecSchemaDefinition"+
				"When shouldUseLegacyTerraformSDKBlockApproachForComplexObjects is called"+
				"Then it should return false with no error", func() {
				p := &specSchemaDefinitionProperty{Type: typeObject, SpecSchemaDefinition: &specSchemaDefinition{}}
				b := p.shouldUseLegacyTerraformSDKBlockApproachForComplexObjects()
				So(b, ShouldBeFalse)
			})

			Convey("Given a specSchemaDefinitionProperty with Type of 'object' and a SpecSchemaDefinition with one specSchemaDefinitionProperty which is blank"+
				"When shouldUseLegacyTerraformSDKBlockApproachForComplexObjects is called"+
				"Then it should return false with no error", func() {
				p := &specSchemaDefinitionProperty{Type: typeObject,
					SpecSchemaDefinition: &specSchemaDefinition{Properties: specSchemaDefinitionProperties{&specSchemaDefinitionProperty{}}}}
				b := p.shouldUseLegacyTerraformSDKBlockApproachForComplexObjects()
				So(b, ShouldBeFalse)
			})
			Convey("Given a specSchemaDefinitionProperty with Type of 'object' and a SpecSchemaDefinition with one specSchemaDefinitionProperty with Type 'boolean' "+
				"When shouldUseLegacyTerraformSDKBlockApproachForComplexObjects is called"+
				"Then it should return true with no error", func() {
				p := &specSchemaDefinitionProperty{Type: typeObject,
					SpecSchemaDefinition: &specSchemaDefinition{Properties: specSchemaDefinitionProperties{&specSchemaDefinitionProperty{Type: typeBool}}}}
				b := p.shouldUseLegacyTerraformSDKBlockApproachForComplexObjects()
				So(b, ShouldBeFalse)
			})
		})

		Convey("returns true", func() {
			Convey("Given a specSchemaDefinitionProperty with Type of 'object' and a SpecSchemaDefinition with one specSchemaDefinitionProperty with Type 'object' "+
				"When shouldUseLegacyTerraformSDKBlockApproachForComplexObjects is called"+
				"Then it should return false with no error", func() {
				p := &specSchemaDefinitionProperty{Type: typeObject,
					SpecSchemaDefinition: &specSchemaDefinition{Properties: specSchemaDefinitionProperties{&specSchemaDefinitionProperty{Type: typeObject}}}}
				b := p.shouldUseLegacyTerraformSDKBlockApproachForComplexObjects()
				So(b, ShouldBeTrue)
			})

			Convey("Given a specSchemaDefinitionProperty with Type of 'object' and a blank SpecSchemaDefinition and EnableLegacyComplexObjectBlockConfiguration = true"+
				"When shouldUseLegacyTerraformSDKBlockApproachForComplexObjects is called"+
				"Then it should return true with no error", func() {
				p := &specSchemaDefinitionProperty{Type: typeObject,
					SpecSchemaDefinition:                        &specSchemaDefinition{},
					EnableLegacyComplexObjectBlockConfiguration: true}
				b := p.shouldUseLegacyTerraformSDKBlockApproachForComplexObjects()
				So(b, ShouldBeTrue)
			})
		})
	})
}
