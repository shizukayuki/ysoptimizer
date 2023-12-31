// Code generated by "enumer --json --type=EquipType equiptype.go"; DO NOT EDIT.

package excel

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _EquipTypeName = "EQUIP_NONEEQUIP_BRACEREQUIP_NECKLACEEQUIP_SHOESEQUIP_RINGEQUIP_DRESSEQUIP_WEAPON"

var _EquipTypeIndex = [...]uint8{0, 10, 22, 36, 47, 57, 68, 80}

const _EquipTypeLowerName = "equip_noneequip_bracerequip_necklaceequip_shoesequip_ringequip_dressequip_weapon"

func (i EquipType) String() string {
	if i >= EquipType(len(_EquipTypeIndex)-1) {
		return fmt.Sprintf("EquipType(%d)", i)
	}
	return _EquipTypeName[_EquipTypeIndex[i]:_EquipTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EquipTypeNoOp() {
	var x [1]struct{}
	_ = x[EQUIP_NONE-(0)]
	_ = x[EQUIP_BRACER-(1)]
	_ = x[EQUIP_NECKLACE-(2)]
	_ = x[EQUIP_SHOES-(3)]
	_ = x[EQUIP_RING-(4)]
	_ = x[EQUIP_DRESS-(5)]
	_ = x[EQUIP_WEAPON-(6)]
}

var _EquipTypeValues = []EquipType{EQUIP_NONE, EQUIP_BRACER, EQUIP_NECKLACE, EQUIP_SHOES, EQUIP_RING, EQUIP_DRESS, EQUIP_WEAPON}

var _EquipTypeNameToValueMap = map[string]EquipType{
	_EquipTypeName[0:10]:       EQUIP_NONE,
	_EquipTypeLowerName[0:10]:  EQUIP_NONE,
	_EquipTypeName[10:22]:      EQUIP_BRACER,
	_EquipTypeLowerName[10:22]: EQUIP_BRACER,
	_EquipTypeName[22:36]:      EQUIP_NECKLACE,
	_EquipTypeLowerName[22:36]: EQUIP_NECKLACE,
	_EquipTypeName[36:47]:      EQUIP_SHOES,
	_EquipTypeLowerName[36:47]: EQUIP_SHOES,
	_EquipTypeName[47:57]:      EQUIP_RING,
	_EquipTypeLowerName[47:57]: EQUIP_RING,
	_EquipTypeName[57:68]:      EQUIP_DRESS,
	_EquipTypeLowerName[57:68]: EQUIP_DRESS,
	_EquipTypeName[68:80]:      EQUIP_WEAPON,
	_EquipTypeLowerName[68:80]: EQUIP_WEAPON,
}

var _EquipTypeNames = []string{
	_EquipTypeName[0:10],
	_EquipTypeName[10:22],
	_EquipTypeName[22:36],
	_EquipTypeName[36:47],
	_EquipTypeName[47:57],
	_EquipTypeName[57:68],
	_EquipTypeName[68:80],
}

// EquipTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EquipTypeString(s string) (EquipType, error) {
	if val, ok := _EquipTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EquipTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to EquipType values", s)
}

// EquipTypeValues returns all values of the enum
func EquipTypeValues() []EquipType {
	return _EquipTypeValues
}

// EquipTypeStrings returns a slice of all String values of the enum
func EquipTypeStrings() []string {
	strs := make([]string, len(_EquipTypeNames))
	copy(strs, _EquipTypeNames)
	return strs
}

// IsAEquipType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i EquipType) IsAEquipType() bool {
	for _, v := range _EquipTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for EquipType
func (i EquipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for EquipType
func (i *EquipType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("EquipType should be a string, got %s", data)
	}

	var err error
	*i, err = EquipTypeString(s)
	return err
}
