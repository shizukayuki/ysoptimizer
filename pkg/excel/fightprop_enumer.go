// Code generated by "enumer --json --type=FightProp fightprop.go"; DO NOT EDIT.

package excel

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _FightPropName = "FIGHT_PROP_NONEFIGHT_PROP_BASE_HPFIGHT_PROP_HPFIGHT_PROP_HP_PERCENTFIGHT_PROP_BASE_ATTACKFIGHT_PROP_ATTACKFIGHT_PROP_ATTACK_PERCENTFIGHT_PROP_BASE_DEFENSEFIGHT_PROP_DEFENSEFIGHT_PROP_DEFENSE_PERCENTFIGHT_PROP_BASE_SPEEDFIGHT_PROP_SPEED_PERCENTFIGHT_PROP_HP_MP_PERCENTFIGHT_PROP_ATTACK_MP_PERCENTFIGHT_PROP_CRITICALFIGHT_PROP_ANTI_CRITICALFIGHT_PROP_CRITICAL_HURTFIGHT_PROP_CHARGE_EFFICIENCYFIGHT_PROP_ADD_HURTFIGHT_PROP_SUB_HURTFIGHT_PROP_HEAL_ADDFIGHT_PROP_HEALED_ADDFIGHT_PROP_ELEMENT_MASTERYFIGHT_PROP_PHYSICAL_SUB_HURTFIGHT_PROP_PHYSICAL_ADD_HURTFIGHT_PROP_DEFENCE_IGNORE_RATIOFIGHT_PROP_DEFENCE_IGNORE_DELTAFIGHT_PROP_FIRE_ADD_HURTFIGHT_PROP_ELEC_ADD_HURTFIGHT_PROP_WATER_ADD_HURTFIGHT_PROP_GRASS_ADD_HURTFIGHT_PROP_WIND_ADD_HURTFIGHT_PROP_ROCK_ADD_HURTFIGHT_PROP_ICE_ADD_HURTFIGHT_PROP_HIT_HEAD_ADD_HURTFIGHT_PROP_FIRE_SUB_HURTFIGHT_PROP_ELEC_SUB_HURTFIGHT_PROP_WATER_SUB_HURTFIGHT_PROP_GRASS_SUB_HURTFIGHT_PROP_WIND_SUB_HURTFIGHT_PROP_ROCK_SUB_HURTFIGHT_PROP_ICE_SUB_HURTFIGHT_PROP_EFFECT_HITFIGHT_PROP_EFFECT_RESISTFIGHT_PROP_FREEZE_RESISTFIGHT_PROP_TORPOR_RESISTFIGHT_PROP_DIZZY_RESISTFIGHT_PROP_FREEZE_SHORTENFIGHT_PROP_TORPOR_SHORTENFIGHT_PROP_DIZZY_SHORTENFIGHT_PROP_MAX_FIRE_ENERGYFIGHT_PROP_MAX_ELEC_ENERGYFIGHT_PROP_MAX_WATER_ENERGYFIGHT_PROP_MAX_GRASS_ENERGYFIGHT_PROP_MAX_WIND_ENERGYFIGHT_PROP_MAX_ICE_ENERGYFIGHT_PROP_MAX_ROCK_ENERGYFIGHT_PROP_SKILL_CD_MINUS_RATIOFIGHT_PROP_SHIELD_COST_MINUS_RATIOFIGHT_PROP_CUR_FIRE_ENERGYFIGHT_PROP_CUR_ELEC_ENERGYFIGHT_PROP_CUR_WATER_ENERGYFIGHT_PROP_CUR_GRASS_ENERGYFIGHT_PROP_CUR_WIND_ENERGYFIGHT_PROP_CUR_ICE_ENERGYFIGHT_PROP_CUR_ROCK_ENERGYFIGHT_PROP_CUR_HPFIGHT_PROP_MAX_HPFIGHT_PROP_CUR_ATTACKFIGHT_PROP_CUR_DEFENSEFIGHT_PROP_CUR_SPEEDFIGHT_PROP_NONEXTRA_ATTACKFIGHT_PROP_NONEXTRA_DEFENSEFIGHT_PROP_NONEXTRA_CRITICALFIGHT_PROP_NONEXTRA_ANTI_CRITICALFIGHT_PROP_NONEXTRA_CRITICAL_HURTFIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCYFIGHT_PROP_NONEXTRA_ELEMENT_MASTERYFIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURTFIGHT_PROP_NONEXTRA_FIRE_ADD_HURTFIGHT_PROP_NONEXTRA_ELEC_ADD_HURTFIGHT_PROP_NONEXTRA_WATER_ADD_HURTFIGHT_PROP_NONEXTRA_GRASS_ADD_HURTFIGHT_PROP_NONEXTRA_WIND_ADD_HURTFIGHT_PROP_NONEXTRA_ROCK_ADD_HURTFIGHT_PROP_NONEXTRA_ICE_ADD_HURTFIGHT_PROP_NONEXTRA_FIRE_SUB_HURTFIGHT_PROP_NONEXTRA_ELEC_SUB_HURTFIGHT_PROP_NONEXTRA_WATER_SUB_HURTFIGHT_PROP_NONEXTRA_GRASS_SUB_HURTFIGHT_PROP_NONEXTRA_WIND_SUB_HURTFIGHT_PROP_NONEXTRA_ROCK_SUB_HURTFIGHT_PROP_NONEXTRA_ICE_SUB_HURTFIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIOFIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIOFIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT"
const _FightPropLowerName = "fight_prop_nonefight_prop_base_hpfight_prop_hpfight_prop_hp_percentfight_prop_base_attackfight_prop_attackfight_prop_attack_percentfight_prop_base_defensefight_prop_defensefight_prop_defense_percentfight_prop_base_speedfight_prop_speed_percentfight_prop_hp_mp_percentfight_prop_attack_mp_percentfight_prop_criticalfight_prop_anti_criticalfight_prop_critical_hurtfight_prop_charge_efficiencyfight_prop_add_hurtfight_prop_sub_hurtfight_prop_heal_addfight_prop_healed_addfight_prop_element_masteryfight_prop_physical_sub_hurtfight_prop_physical_add_hurtfight_prop_defence_ignore_ratiofight_prop_defence_ignore_deltafight_prop_fire_add_hurtfight_prop_elec_add_hurtfight_prop_water_add_hurtfight_prop_grass_add_hurtfight_prop_wind_add_hurtfight_prop_rock_add_hurtfight_prop_ice_add_hurtfight_prop_hit_head_add_hurtfight_prop_fire_sub_hurtfight_prop_elec_sub_hurtfight_prop_water_sub_hurtfight_prop_grass_sub_hurtfight_prop_wind_sub_hurtfight_prop_rock_sub_hurtfight_prop_ice_sub_hurtfight_prop_effect_hitfight_prop_effect_resistfight_prop_freeze_resistfight_prop_torpor_resistfight_prop_dizzy_resistfight_prop_freeze_shortenfight_prop_torpor_shortenfight_prop_dizzy_shortenfight_prop_max_fire_energyfight_prop_max_elec_energyfight_prop_max_water_energyfight_prop_max_grass_energyfight_prop_max_wind_energyfight_prop_max_ice_energyfight_prop_max_rock_energyfight_prop_skill_cd_minus_ratiofight_prop_shield_cost_minus_ratiofight_prop_cur_fire_energyfight_prop_cur_elec_energyfight_prop_cur_water_energyfight_prop_cur_grass_energyfight_prop_cur_wind_energyfight_prop_cur_ice_energyfight_prop_cur_rock_energyfight_prop_cur_hpfight_prop_max_hpfight_prop_cur_attackfight_prop_cur_defensefight_prop_cur_speedfight_prop_nonextra_attackfight_prop_nonextra_defensefight_prop_nonextra_criticalfight_prop_nonextra_anti_criticalfight_prop_nonextra_critical_hurtfight_prop_nonextra_charge_efficiencyfight_prop_nonextra_element_masteryfight_prop_nonextra_physical_sub_hurtfight_prop_nonextra_fire_add_hurtfight_prop_nonextra_elec_add_hurtfight_prop_nonextra_water_add_hurtfight_prop_nonextra_grass_add_hurtfight_prop_nonextra_wind_add_hurtfight_prop_nonextra_rock_add_hurtfight_prop_nonextra_ice_add_hurtfight_prop_nonextra_fire_sub_hurtfight_prop_nonextra_elec_sub_hurtfight_prop_nonextra_water_sub_hurtfight_prop_nonextra_grass_sub_hurtfight_prop_nonextra_wind_sub_hurtfight_prop_nonextra_rock_sub_hurtfight_prop_nonextra_ice_sub_hurtfight_prop_nonextra_skill_cd_minus_ratiofight_prop_nonextra_shield_cost_minus_ratiofight_prop_nonextra_physical_add_hurt"

var _FightPropMap = map[FightProp]string{
	0:    _FightPropName[0:15],
	1:    _FightPropName[15:33],
	2:    _FightPropName[33:46],
	3:    _FightPropName[46:67],
	4:    _FightPropName[67:89],
	5:    _FightPropName[89:106],
	6:    _FightPropName[106:131],
	7:    _FightPropName[131:154],
	8:    _FightPropName[154:172],
	9:    _FightPropName[172:198],
	10:   _FightPropName[198:219],
	11:   _FightPropName[219:243],
	12:   _FightPropName[243:267],
	13:   _FightPropName[267:295],
	20:   _FightPropName[295:314],
	21:   _FightPropName[314:338],
	22:   _FightPropName[338:362],
	23:   _FightPropName[362:390],
	24:   _FightPropName[390:409],
	25:   _FightPropName[409:428],
	26:   _FightPropName[428:447],
	27:   _FightPropName[447:468],
	28:   _FightPropName[468:494],
	29:   _FightPropName[494:522],
	30:   _FightPropName[522:550],
	31:   _FightPropName[550:581],
	32:   _FightPropName[581:612],
	40:   _FightPropName[612:636],
	41:   _FightPropName[636:660],
	42:   _FightPropName[660:685],
	43:   _FightPropName[685:710],
	44:   _FightPropName[710:734],
	45:   _FightPropName[734:758],
	46:   _FightPropName[758:781],
	47:   _FightPropName[781:809],
	50:   _FightPropName[809:833],
	51:   _FightPropName[833:857],
	52:   _FightPropName[857:882],
	53:   _FightPropName[882:907],
	54:   _FightPropName[907:931],
	55:   _FightPropName[931:955],
	56:   _FightPropName[955:978],
	60:   _FightPropName[978:999],
	61:   _FightPropName[999:1023],
	62:   _FightPropName[1023:1047],
	63:   _FightPropName[1047:1071],
	64:   _FightPropName[1071:1094],
	65:   _FightPropName[1094:1119],
	66:   _FightPropName[1119:1144],
	67:   _FightPropName[1144:1168],
	70:   _FightPropName[1168:1194],
	71:   _FightPropName[1194:1220],
	72:   _FightPropName[1220:1247],
	73:   _FightPropName[1247:1274],
	74:   _FightPropName[1274:1300],
	75:   _FightPropName[1300:1325],
	76:   _FightPropName[1325:1351],
	80:   _FightPropName[1351:1382],
	81:   _FightPropName[1382:1416],
	1000: _FightPropName[1416:1442],
	1001: _FightPropName[1442:1468],
	1002: _FightPropName[1468:1495],
	1003: _FightPropName[1495:1522],
	1004: _FightPropName[1522:1548],
	1005: _FightPropName[1548:1573],
	1006: _FightPropName[1573:1599],
	1010: _FightPropName[1599:1616],
	2000: _FightPropName[1616:1633],
	2001: _FightPropName[1633:1654],
	2002: _FightPropName[1654:1676],
	2003: _FightPropName[1676:1696],
	3000: _FightPropName[1696:1722],
	3001: _FightPropName[1722:1749],
	3002: _FightPropName[1749:1777],
	3003: _FightPropName[1777:1810],
	3004: _FightPropName[1810:1843],
	3005: _FightPropName[1843:1880],
	3006: _FightPropName[1880:1915],
	3007: _FightPropName[1915:1952],
	3008: _FightPropName[1952:1985],
	3009: _FightPropName[1985:2018],
	3010: _FightPropName[2018:2052],
	3011: _FightPropName[2052:2086],
	3012: _FightPropName[2086:2119],
	3013: _FightPropName[2119:2152],
	3014: _FightPropName[2152:2184],
	3015: _FightPropName[2184:2217],
	3016: _FightPropName[2217:2250],
	3017: _FightPropName[2250:2284],
	3018: _FightPropName[2284:2318],
	3019: _FightPropName[2318:2351],
	3020: _FightPropName[2351:2384],
	3021: _FightPropName[2384:2416],
	3022: _FightPropName[2416:2456],
	3023: _FightPropName[2456:2499],
	3024: _FightPropName[2499:2536],
}

func (i FightProp) String() string {
	if str, ok := _FightPropMap[i]; ok {
		return str
	}
	return fmt.Sprintf("FightProp(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FightPropNoOp() {
	var x [1]struct{}
	_ = x[FIGHT_PROP_NONE-(0)]
	_ = x[FIGHT_PROP_BASE_HP-(1)]
	_ = x[FIGHT_PROP_HP-(2)]
	_ = x[FIGHT_PROP_HP_PERCENT-(3)]
	_ = x[FIGHT_PROP_BASE_ATTACK-(4)]
	_ = x[FIGHT_PROP_ATTACK-(5)]
	_ = x[FIGHT_PROP_ATTACK_PERCENT-(6)]
	_ = x[FIGHT_PROP_BASE_DEFENSE-(7)]
	_ = x[FIGHT_PROP_DEFENSE-(8)]
	_ = x[FIGHT_PROP_DEFENSE_PERCENT-(9)]
	_ = x[FIGHT_PROP_BASE_SPEED-(10)]
	_ = x[FIGHT_PROP_SPEED_PERCENT-(11)]
	_ = x[FIGHT_PROP_HP_MP_PERCENT-(12)]
	_ = x[FIGHT_PROP_ATTACK_MP_PERCENT-(13)]
	_ = x[FIGHT_PROP_CRITICAL-(20)]
	_ = x[FIGHT_PROP_ANTI_CRITICAL-(21)]
	_ = x[FIGHT_PROP_CRITICAL_HURT-(22)]
	_ = x[FIGHT_PROP_CHARGE_EFFICIENCY-(23)]
	_ = x[FIGHT_PROP_ADD_HURT-(24)]
	_ = x[FIGHT_PROP_SUB_HURT-(25)]
	_ = x[FIGHT_PROP_HEAL_ADD-(26)]
	_ = x[FIGHT_PROP_HEALED_ADD-(27)]
	_ = x[FIGHT_PROP_ELEMENT_MASTERY-(28)]
	_ = x[FIGHT_PROP_PHYSICAL_SUB_HURT-(29)]
	_ = x[FIGHT_PROP_PHYSICAL_ADD_HURT-(30)]
	_ = x[FIGHT_PROP_DEFENCE_IGNORE_RATIO-(31)]
	_ = x[FIGHT_PROP_DEFENCE_IGNORE_DELTA-(32)]
	_ = x[FIGHT_PROP_FIRE_ADD_HURT-(40)]
	_ = x[FIGHT_PROP_ELEC_ADD_HURT-(41)]
	_ = x[FIGHT_PROP_WATER_ADD_HURT-(42)]
	_ = x[FIGHT_PROP_GRASS_ADD_HURT-(43)]
	_ = x[FIGHT_PROP_WIND_ADD_HURT-(44)]
	_ = x[FIGHT_PROP_ROCK_ADD_HURT-(45)]
	_ = x[FIGHT_PROP_ICE_ADD_HURT-(46)]
	_ = x[FIGHT_PROP_HIT_HEAD_ADD_HURT-(47)]
	_ = x[FIGHT_PROP_FIRE_SUB_HURT-(50)]
	_ = x[FIGHT_PROP_ELEC_SUB_HURT-(51)]
	_ = x[FIGHT_PROP_WATER_SUB_HURT-(52)]
	_ = x[FIGHT_PROP_GRASS_SUB_HURT-(53)]
	_ = x[FIGHT_PROP_WIND_SUB_HURT-(54)]
	_ = x[FIGHT_PROP_ROCK_SUB_HURT-(55)]
	_ = x[FIGHT_PROP_ICE_SUB_HURT-(56)]
	_ = x[FIGHT_PROP_EFFECT_HIT-(60)]
	_ = x[FIGHT_PROP_EFFECT_RESIST-(61)]
	_ = x[FIGHT_PROP_FREEZE_RESIST-(62)]
	_ = x[FIGHT_PROP_TORPOR_RESIST-(63)]
	_ = x[FIGHT_PROP_DIZZY_RESIST-(64)]
	_ = x[FIGHT_PROP_FREEZE_SHORTEN-(65)]
	_ = x[FIGHT_PROP_TORPOR_SHORTEN-(66)]
	_ = x[FIGHT_PROP_DIZZY_SHORTEN-(67)]
	_ = x[FIGHT_PROP_MAX_FIRE_ENERGY-(70)]
	_ = x[FIGHT_PROP_MAX_ELEC_ENERGY-(71)]
	_ = x[FIGHT_PROP_MAX_WATER_ENERGY-(72)]
	_ = x[FIGHT_PROP_MAX_GRASS_ENERGY-(73)]
	_ = x[FIGHT_PROP_MAX_WIND_ENERGY-(74)]
	_ = x[FIGHT_PROP_MAX_ICE_ENERGY-(75)]
	_ = x[FIGHT_PROP_MAX_ROCK_ENERGY-(76)]
	_ = x[FIGHT_PROP_SKILL_CD_MINUS_RATIO-(80)]
	_ = x[FIGHT_PROP_SHIELD_COST_MINUS_RATIO-(81)]
	_ = x[FIGHT_PROP_CUR_FIRE_ENERGY-(1000)]
	_ = x[FIGHT_PROP_CUR_ELEC_ENERGY-(1001)]
	_ = x[FIGHT_PROP_CUR_WATER_ENERGY-(1002)]
	_ = x[FIGHT_PROP_CUR_GRASS_ENERGY-(1003)]
	_ = x[FIGHT_PROP_CUR_WIND_ENERGY-(1004)]
	_ = x[FIGHT_PROP_CUR_ICE_ENERGY-(1005)]
	_ = x[FIGHT_PROP_CUR_ROCK_ENERGY-(1006)]
	_ = x[FIGHT_PROP_CUR_HP-(1010)]
	_ = x[FIGHT_PROP_MAX_HP-(2000)]
	_ = x[FIGHT_PROP_CUR_ATTACK-(2001)]
	_ = x[FIGHT_PROP_CUR_DEFENSE-(2002)]
	_ = x[FIGHT_PROP_CUR_SPEED-(2003)]
	_ = x[FIGHT_PROP_NONEXTRA_ATTACK-(3000)]
	_ = x[FIGHT_PROP_NONEXTRA_DEFENSE-(3001)]
	_ = x[FIGHT_PROP_NONEXTRA_CRITICAL-(3002)]
	_ = x[FIGHT_PROP_NONEXTRA_ANTI_CRITICAL-(3003)]
	_ = x[FIGHT_PROP_NONEXTRA_CRITICAL_HURT-(3004)]
	_ = x[FIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCY-(3005)]
	_ = x[FIGHT_PROP_NONEXTRA_ELEMENT_MASTERY-(3006)]
	_ = x[FIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURT-(3007)]
	_ = x[FIGHT_PROP_NONEXTRA_FIRE_ADD_HURT-(3008)]
	_ = x[FIGHT_PROP_NONEXTRA_ELEC_ADD_HURT-(3009)]
	_ = x[FIGHT_PROP_NONEXTRA_WATER_ADD_HURT-(3010)]
	_ = x[FIGHT_PROP_NONEXTRA_GRASS_ADD_HURT-(3011)]
	_ = x[FIGHT_PROP_NONEXTRA_WIND_ADD_HURT-(3012)]
	_ = x[FIGHT_PROP_NONEXTRA_ROCK_ADD_HURT-(3013)]
	_ = x[FIGHT_PROP_NONEXTRA_ICE_ADD_HURT-(3014)]
	_ = x[FIGHT_PROP_NONEXTRA_FIRE_SUB_HURT-(3015)]
	_ = x[FIGHT_PROP_NONEXTRA_ELEC_SUB_HURT-(3016)]
	_ = x[FIGHT_PROP_NONEXTRA_WATER_SUB_HURT-(3017)]
	_ = x[FIGHT_PROP_NONEXTRA_GRASS_SUB_HURT-(3018)]
	_ = x[FIGHT_PROP_NONEXTRA_WIND_SUB_HURT-(3019)]
	_ = x[FIGHT_PROP_NONEXTRA_ROCK_SUB_HURT-(3020)]
	_ = x[FIGHT_PROP_NONEXTRA_ICE_SUB_HURT-(3021)]
	_ = x[FIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIO-(3022)]
	_ = x[FIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIO-(3023)]
	_ = x[FIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT-(3024)]
}

var _FightPropValues = []FightProp{FIGHT_PROP_NONE, FIGHT_PROP_BASE_HP, FIGHT_PROP_HP, FIGHT_PROP_HP_PERCENT, FIGHT_PROP_BASE_ATTACK, FIGHT_PROP_ATTACK, FIGHT_PROP_ATTACK_PERCENT, FIGHT_PROP_BASE_DEFENSE, FIGHT_PROP_DEFENSE, FIGHT_PROP_DEFENSE_PERCENT, FIGHT_PROP_BASE_SPEED, FIGHT_PROP_SPEED_PERCENT, FIGHT_PROP_HP_MP_PERCENT, FIGHT_PROP_ATTACK_MP_PERCENT, FIGHT_PROP_CRITICAL, FIGHT_PROP_ANTI_CRITICAL, FIGHT_PROP_CRITICAL_HURT, FIGHT_PROP_CHARGE_EFFICIENCY, FIGHT_PROP_ADD_HURT, FIGHT_PROP_SUB_HURT, FIGHT_PROP_HEAL_ADD, FIGHT_PROP_HEALED_ADD, FIGHT_PROP_ELEMENT_MASTERY, FIGHT_PROP_PHYSICAL_SUB_HURT, FIGHT_PROP_PHYSICAL_ADD_HURT, FIGHT_PROP_DEFENCE_IGNORE_RATIO, FIGHT_PROP_DEFENCE_IGNORE_DELTA, FIGHT_PROP_FIRE_ADD_HURT, FIGHT_PROP_ELEC_ADD_HURT, FIGHT_PROP_WATER_ADD_HURT, FIGHT_PROP_GRASS_ADD_HURT, FIGHT_PROP_WIND_ADD_HURT, FIGHT_PROP_ROCK_ADD_HURT, FIGHT_PROP_ICE_ADD_HURT, FIGHT_PROP_HIT_HEAD_ADD_HURT, FIGHT_PROP_FIRE_SUB_HURT, FIGHT_PROP_ELEC_SUB_HURT, FIGHT_PROP_WATER_SUB_HURT, FIGHT_PROP_GRASS_SUB_HURT, FIGHT_PROP_WIND_SUB_HURT, FIGHT_PROP_ROCK_SUB_HURT, FIGHT_PROP_ICE_SUB_HURT, FIGHT_PROP_EFFECT_HIT, FIGHT_PROP_EFFECT_RESIST, FIGHT_PROP_FREEZE_RESIST, FIGHT_PROP_TORPOR_RESIST, FIGHT_PROP_DIZZY_RESIST, FIGHT_PROP_FREEZE_SHORTEN, FIGHT_PROP_TORPOR_SHORTEN, FIGHT_PROP_DIZZY_SHORTEN, FIGHT_PROP_MAX_FIRE_ENERGY, FIGHT_PROP_MAX_ELEC_ENERGY, FIGHT_PROP_MAX_WATER_ENERGY, FIGHT_PROP_MAX_GRASS_ENERGY, FIGHT_PROP_MAX_WIND_ENERGY, FIGHT_PROP_MAX_ICE_ENERGY, FIGHT_PROP_MAX_ROCK_ENERGY, FIGHT_PROP_SKILL_CD_MINUS_RATIO, FIGHT_PROP_SHIELD_COST_MINUS_RATIO, FIGHT_PROP_CUR_FIRE_ENERGY, FIGHT_PROP_CUR_ELEC_ENERGY, FIGHT_PROP_CUR_WATER_ENERGY, FIGHT_PROP_CUR_GRASS_ENERGY, FIGHT_PROP_CUR_WIND_ENERGY, FIGHT_PROP_CUR_ICE_ENERGY, FIGHT_PROP_CUR_ROCK_ENERGY, FIGHT_PROP_CUR_HP, FIGHT_PROP_MAX_HP, FIGHT_PROP_CUR_ATTACK, FIGHT_PROP_CUR_DEFENSE, FIGHT_PROP_CUR_SPEED, FIGHT_PROP_NONEXTRA_ATTACK, FIGHT_PROP_NONEXTRA_DEFENSE, FIGHT_PROP_NONEXTRA_CRITICAL, FIGHT_PROP_NONEXTRA_ANTI_CRITICAL, FIGHT_PROP_NONEXTRA_CRITICAL_HURT, FIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCY, FIGHT_PROP_NONEXTRA_ELEMENT_MASTERY, FIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURT, FIGHT_PROP_NONEXTRA_FIRE_ADD_HURT, FIGHT_PROP_NONEXTRA_ELEC_ADD_HURT, FIGHT_PROP_NONEXTRA_WATER_ADD_HURT, FIGHT_PROP_NONEXTRA_GRASS_ADD_HURT, FIGHT_PROP_NONEXTRA_WIND_ADD_HURT, FIGHT_PROP_NONEXTRA_ROCK_ADD_HURT, FIGHT_PROP_NONEXTRA_ICE_ADD_HURT, FIGHT_PROP_NONEXTRA_FIRE_SUB_HURT, FIGHT_PROP_NONEXTRA_ELEC_SUB_HURT, FIGHT_PROP_NONEXTRA_WATER_SUB_HURT, FIGHT_PROP_NONEXTRA_GRASS_SUB_HURT, FIGHT_PROP_NONEXTRA_WIND_SUB_HURT, FIGHT_PROP_NONEXTRA_ROCK_SUB_HURT, FIGHT_PROP_NONEXTRA_ICE_SUB_HURT, FIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIO, FIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIO, FIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT}

var _FightPropNameToValueMap = map[string]FightProp{
	_FightPropName[0:15]:           FIGHT_PROP_NONE,
	_FightPropLowerName[0:15]:      FIGHT_PROP_NONE,
	_FightPropName[15:33]:          FIGHT_PROP_BASE_HP,
	_FightPropLowerName[15:33]:     FIGHT_PROP_BASE_HP,
	_FightPropName[33:46]:          FIGHT_PROP_HP,
	_FightPropLowerName[33:46]:     FIGHT_PROP_HP,
	_FightPropName[46:67]:          FIGHT_PROP_HP_PERCENT,
	_FightPropLowerName[46:67]:     FIGHT_PROP_HP_PERCENT,
	_FightPropName[67:89]:          FIGHT_PROP_BASE_ATTACK,
	_FightPropLowerName[67:89]:     FIGHT_PROP_BASE_ATTACK,
	_FightPropName[89:106]:         FIGHT_PROP_ATTACK,
	_FightPropLowerName[89:106]:    FIGHT_PROP_ATTACK,
	_FightPropName[106:131]:        FIGHT_PROP_ATTACK_PERCENT,
	_FightPropLowerName[106:131]:   FIGHT_PROP_ATTACK_PERCENT,
	_FightPropName[131:154]:        FIGHT_PROP_BASE_DEFENSE,
	_FightPropLowerName[131:154]:   FIGHT_PROP_BASE_DEFENSE,
	_FightPropName[154:172]:        FIGHT_PROP_DEFENSE,
	_FightPropLowerName[154:172]:   FIGHT_PROP_DEFENSE,
	_FightPropName[172:198]:        FIGHT_PROP_DEFENSE_PERCENT,
	_FightPropLowerName[172:198]:   FIGHT_PROP_DEFENSE_PERCENT,
	_FightPropName[198:219]:        FIGHT_PROP_BASE_SPEED,
	_FightPropLowerName[198:219]:   FIGHT_PROP_BASE_SPEED,
	_FightPropName[219:243]:        FIGHT_PROP_SPEED_PERCENT,
	_FightPropLowerName[219:243]:   FIGHT_PROP_SPEED_PERCENT,
	_FightPropName[243:267]:        FIGHT_PROP_HP_MP_PERCENT,
	_FightPropLowerName[243:267]:   FIGHT_PROP_HP_MP_PERCENT,
	_FightPropName[267:295]:        FIGHT_PROP_ATTACK_MP_PERCENT,
	_FightPropLowerName[267:295]:   FIGHT_PROP_ATTACK_MP_PERCENT,
	_FightPropName[295:314]:        FIGHT_PROP_CRITICAL,
	_FightPropLowerName[295:314]:   FIGHT_PROP_CRITICAL,
	_FightPropName[314:338]:        FIGHT_PROP_ANTI_CRITICAL,
	_FightPropLowerName[314:338]:   FIGHT_PROP_ANTI_CRITICAL,
	_FightPropName[338:362]:        FIGHT_PROP_CRITICAL_HURT,
	_FightPropLowerName[338:362]:   FIGHT_PROP_CRITICAL_HURT,
	_FightPropName[362:390]:        FIGHT_PROP_CHARGE_EFFICIENCY,
	_FightPropLowerName[362:390]:   FIGHT_PROP_CHARGE_EFFICIENCY,
	_FightPropName[390:409]:        FIGHT_PROP_ADD_HURT,
	_FightPropLowerName[390:409]:   FIGHT_PROP_ADD_HURT,
	_FightPropName[409:428]:        FIGHT_PROP_SUB_HURT,
	_FightPropLowerName[409:428]:   FIGHT_PROP_SUB_HURT,
	_FightPropName[428:447]:        FIGHT_PROP_HEAL_ADD,
	_FightPropLowerName[428:447]:   FIGHT_PROP_HEAL_ADD,
	_FightPropName[447:468]:        FIGHT_PROP_HEALED_ADD,
	_FightPropLowerName[447:468]:   FIGHT_PROP_HEALED_ADD,
	_FightPropName[468:494]:        FIGHT_PROP_ELEMENT_MASTERY,
	_FightPropLowerName[468:494]:   FIGHT_PROP_ELEMENT_MASTERY,
	_FightPropName[494:522]:        FIGHT_PROP_PHYSICAL_SUB_HURT,
	_FightPropLowerName[494:522]:   FIGHT_PROP_PHYSICAL_SUB_HURT,
	_FightPropName[522:550]:        FIGHT_PROP_PHYSICAL_ADD_HURT,
	_FightPropLowerName[522:550]:   FIGHT_PROP_PHYSICAL_ADD_HURT,
	_FightPropName[550:581]:        FIGHT_PROP_DEFENCE_IGNORE_RATIO,
	_FightPropLowerName[550:581]:   FIGHT_PROP_DEFENCE_IGNORE_RATIO,
	_FightPropName[581:612]:        FIGHT_PROP_DEFENCE_IGNORE_DELTA,
	_FightPropLowerName[581:612]:   FIGHT_PROP_DEFENCE_IGNORE_DELTA,
	_FightPropName[612:636]:        FIGHT_PROP_FIRE_ADD_HURT,
	_FightPropLowerName[612:636]:   FIGHT_PROP_FIRE_ADD_HURT,
	_FightPropName[636:660]:        FIGHT_PROP_ELEC_ADD_HURT,
	_FightPropLowerName[636:660]:   FIGHT_PROP_ELEC_ADD_HURT,
	_FightPropName[660:685]:        FIGHT_PROP_WATER_ADD_HURT,
	_FightPropLowerName[660:685]:   FIGHT_PROP_WATER_ADD_HURT,
	_FightPropName[685:710]:        FIGHT_PROP_GRASS_ADD_HURT,
	_FightPropLowerName[685:710]:   FIGHT_PROP_GRASS_ADD_HURT,
	_FightPropName[710:734]:        FIGHT_PROP_WIND_ADD_HURT,
	_FightPropLowerName[710:734]:   FIGHT_PROP_WIND_ADD_HURT,
	_FightPropName[734:758]:        FIGHT_PROP_ROCK_ADD_HURT,
	_FightPropLowerName[734:758]:   FIGHT_PROP_ROCK_ADD_HURT,
	_FightPropName[758:781]:        FIGHT_PROP_ICE_ADD_HURT,
	_FightPropLowerName[758:781]:   FIGHT_PROP_ICE_ADD_HURT,
	_FightPropName[781:809]:        FIGHT_PROP_HIT_HEAD_ADD_HURT,
	_FightPropLowerName[781:809]:   FIGHT_PROP_HIT_HEAD_ADD_HURT,
	_FightPropName[809:833]:        FIGHT_PROP_FIRE_SUB_HURT,
	_FightPropLowerName[809:833]:   FIGHT_PROP_FIRE_SUB_HURT,
	_FightPropName[833:857]:        FIGHT_PROP_ELEC_SUB_HURT,
	_FightPropLowerName[833:857]:   FIGHT_PROP_ELEC_SUB_HURT,
	_FightPropName[857:882]:        FIGHT_PROP_WATER_SUB_HURT,
	_FightPropLowerName[857:882]:   FIGHT_PROP_WATER_SUB_HURT,
	_FightPropName[882:907]:        FIGHT_PROP_GRASS_SUB_HURT,
	_FightPropLowerName[882:907]:   FIGHT_PROP_GRASS_SUB_HURT,
	_FightPropName[907:931]:        FIGHT_PROP_WIND_SUB_HURT,
	_FightPropLowerName[907:931]:   FIGHT_PROP_WIND_SUB_HURT,
	_FightPropName[931:955]:        FIGHT_PROP_ROCK_SUB_HURT,
	_FightPropLowerName[931:955]:   FIGHT_PROP_ROCK_SUB_HURT,
	_FightPropName[955:978]:        FIGHT_PROP_ICE_SUB_HURT,
	_FightPropLowerName[955:978]:   FIGHT_PROP_ICE_SUB_HURT,
	_FightPropName[978:999]:        FIGHT_PROP_EFFECT_HIT,
	_FightPropLowerName[978:999]:   FIGHT_PROP_EFFECT_HIT,
	_FightPropName[999:1023]:       FIGHT_PROP_EFFECT_RESIST,
	_FightPropLowerName[999:1023]:  FIGHT_PROP_EFFECT_RESIST,
	_FightPropName[1023:1047]:      FIGHT_PROP_FREEZE_RESIST,
	_FightPropLowerName[1023:1047]: FIGHT_PROP_FREEZE_RESIST,
	_FightPropName[1047:1071]:      FIGHT_PROP_TORPOR_RESIST,
	_FightPropLowerName[1047:1071]: FIGHT_PROP_TORPOR_RESIST,
	_FightPropName[1071:1094]:      FIGHT_PROP_DIZZY_RESIST,
	_FightPropLowerName[1071:1094]: FIGHT_PROP_DIZZY_RESIST,
	_FightPropName[1094:1119]:      FIGHT_PROP_FREEZE_SHORTEN,
	_FightPropLowerName[1094:1119]: FIGHT_PROP_FREEZE_SHORTEN,
	_FightPropName[1119:1144]:      FIGHT_PROP_TORPOR_SHORTEN,
	_FightPropLowerName[1119:1144]: FIGHT_PROP_TORPOR_SHORTEN,
	_FightPropName[1144:1168]:      FIGHT_PROP_DIZZY_SHORTEN,
	_FightPropLowerName[1144:1168]: FIGHT_PROP_DIZZY_SHORTEN,
	_FightPropName[1168:1194]:      FIGHT_PROP_MAX_FIRE_ENERGY,
	_FightPropLowerName[1168:1194]: FIGHT_PROP_MAX_FIRE_ENERGY,
	_FightPropName[1194:1220]:      FIGHT_PROP_MAX_ELEC_ENERGY,
	_FightPropLowerName[1194:1220]: FIGHT_PROP_MAX_ELEC_ENERGY,
	_FightPropName[1220:1247]:      FIGHT_PROP_MAX_WATER_ENERGY,
	_FightPropLowerName[1220:1247]: FIGHT_PROP_MAX_WATER_ENERGY,
	_FightPropName[1247:1274]:      FIGHT_PROP_MAX_GRASS_ENERGY,
	_FightPropLowerName[1247:1274]: FIGHT_PROP_MAX_GRASS_ENERGY,
	_FightPropName[1274:1300]:      FIGHT_PROP_MAX_WIND_ENERGY,
	_FightPropLowerName[1274:1300]: FIGHT_PROP_MAX_WIND_ENERGY,
	_FightPropName[1300:1325]:      FIGHT_PROP_MAX_ICE_ENERGY,
	_FightPropLowerName[1300:1325]: FIGHT_PROP_MAX_ICE_ENERGY,
	_FightPropName[1325:1351]:      FIGHT_PROP_MAX_ROCK_ENERGY,
	_FightPropLowerName[1325:1351]: FIGHT_PROP_MAX_ROCK_ENERGY,
	_FightPropName[1351:1382]:      FIGHT_PROP_SKILL_CD_MINUS_RATIO,
	_FightPropLowerName[1351:1382]: FIGHT_PROP_SKILL_CD_MINUS_RATIO,
	_FightPropName[1382:1416]:      FIGHT_PROP_SHIELD_COST_MINUS_RATIO,
	_FightPropLowerName[1382:1416]: FIGHT_PROP_SHIELD_COST_MINUS_RATIO,
	_FightPropName[1416:1442]:      FIGHT_PROP_CUR_FIRE_ENERGY,
	_FightPropLowerName[1416:1442]: FIGHT_PROP_CUR_FIRE_ENERGY,
	_FightPropName[1442:1468]:      FIGHT_PROP_CUR_ELEC_ENERGY,
	_FightPropLowerName[1442:1468]: FIGHT_PROP_CUR_ELEC_ENERGY,
	_FightPropName[1468:1495]:      FIGHT_PROP_CUR_WATER_ENERGY,
	_FightPropLowerName[1468:1495]: FIGHT_PROP_CUR_WATER_ENERGY,
	_FightPropName[1495:1522]:      FIGHT_PROP_CUR_GRASS_ENERGY,
	_FightPropLowerName[1495:1522]: FIGHT_PROP_CUR_GRASS_ENERGY,
	_FightPropName[1522:1548]:      FIGHT_PROP_CUR_WIND_ENERGY,
	_FightPropLowerName[1522:1548]: FIGHT_PROP_CUR_WIND_ENERGY,
	_FightPropName[1548:1573]:      FIGHT_PROP_CUR_ICE_ENERGY,
	_FightPropLowerName[1548:1573]: FIGHT_PROP_CUR_ICE_ENERGY,
	_FightPropName[1573:1599]:      FIGHT_PROP_CUR_ROCK_ENERGY,
	_FightPropLowerName[1573:1599]: FIGHT_PROP_CUR_ROCK_ENERGY,
	_FightPropName[1599:1616]:      FIGHT_PROP_CUR_HP,
	_FightPropLowerName[1599:1616]: FIGHT_PROP_CUR_HP,
	_FightPropName[1616:1633]:      FIGHT_PROP_MAX_HP,
	_FightPropLowerName[1616:1633]: FIGHT_PROP_MAX_HP,
	_FightPropName[1633:1654]:      FIGHT_PROP_CUR_ATTACK,
	_FightPropLowerName[1633:1654]: FIGHT_PROP_CUR_ATTACK,
	_FightPropName[1654:1676]:      FIGHT_PROP_CUR_DEFENSE,
	_FightPropLowerName[1654:1676]: FIGHT_PROP_CUR_DEFENSE,
	_FightPropName[1676:1696]:      FIGHT_PROP_CUR_SPEED,
	_FightPropLowerName[1676:1696]: FIGHT_PROP_CUR_SPEED,
	_FightPropName[1696:1722]:      FIGHT_PROP_NONEXTRA_ATTACK,
	_FightPropLowerName[1696:1722]: FIGHT_PROP_NONEXTRA_ATTACK,
	_FightPropName[1722:1749]:      FIGHT_PROP_NONEXTRA_DEFENSE,
	_FightPropLowerName[1722:1749]: FIGHT_PROP_NONEXTRA_DEFENSE,
	_FightPropName[1749:1777]:      FIGHT_PROP_NONEXTRA_CRITICAL,
	_FightPropLowerName[1749:1777]: FIGHT_PROP_NONEXTRA_CRITICAL,
	_FightPropName[1777:1810]:      FIGHT_PROP_NONEXTRA_ANTI_CRITICAL,
	_FightPropLowerName[1777:1810]: FIGHT_PROP_NONEXTRA_ANTI_CRITICAL,
	_FightPropName[1810:1843]:      FIGHT_PROP_NONEXTRA_CRITICAL_HURT,
	_FightPropLowerName[1810:1843]: FIGHT_PROP_NONEXTRA_CRITICAL_HURT,
	_FightPropName[1843:1880]:      FIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCY,
	_FightPropLowerName[1843:1880]: FIGHT_PROP_NONEXTRA_CHARGE_EFFICIENCY,
	_FightPropName[1880:1915]:      FIGHT_PROP_NONEXTRA_ELEMENT_MASTERY,
	_FightPropLowerName[1880:1915]: FIGHT_PROP_NONEXTRA_ELEMENT_MASTERY,
	_FightPropName[1915:1952]:      FIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURT,
	_FightPropLowerName[1915:1952]: FIGHT_PROP_NONEXTRA_PHYSICAL_SUB_HURT,
	_FightPropName[1952:1985]:      FIGHT_PROP_NONEXTRA_FIRE_ADD_HURT,
	_FightPropLowerName[1952:1985]: FIGHT_PROP_NONEXTRA_FIRE_ADD_HURT,
	_FightPropName[1985:2018]:      FIGHT_PROP_NONEXTRA_ELEC_ADD_HURT,
	_FightPropLowerName[1985:2018]: FIGHT_PROP_NONEXTRA_ELEC_ADD_HURT,
	_FightPropName[2018:2052]:      FIGHT_PROP_NONEXTRA_WATER_ADD_HURT,
	_FightPropLowerName[2018:2052]: FIGHT_PROP_NONEXTRA_WATER_ADD_HURT,
	_FightPropName[2052:2086]:      FIGHT_PROP_NONEXTRA_GRASS_ADD_HURT,
	_FightPropLowerName[2052:2086]: FIGHT_PROP_NONEXTRA_GRASS_ADD_HURT,
	_FightPropName[2086:2119]:      FIGHT_PROP_NONEXTRA_WIND_ADD_HURT,
	_FightPropLowerName[2086:2119]: FIGHT_PROP_NONEXTRA_WIND_ADD_HURT,
	_FightPropName[2119:2152]:      FIGHT_PROP_NONEXTRA_ROCK_ADD_HURT,
	_FightPropLowerName[2119:2152]: FIGHT_PROP_NONEXTRA_ROCK_ADD_HURT,
	_FightPropName[2152:2184]:      FIGHT_PROP_NONEXTRA_ICE_ADD_HURT,
	_FightPropLowerName[2152:2184]: FIGHT_PROP_NONEXTRA_ICE_ADD_HURT,
	_FightPropName[2184:2217]:      FIGHT_PROP_NONEXTRA_FIRE_SUB_HURT,
	_FightPropLowerName[2184:2217]: FIGHT_PROP_NONEXTRA_FIRE_SUB_HURT,
	_FightPropName[2217:2250]:      FIGHT_PROP_NONEXTRA_ELEC_SUB_HURT,
	_FightPropLowerName[2217:2250]: FIGHT_PROP_NONEXTRA_ELEC_SUB_HURT,
	_FightPropName[2250:2284]:      FIGHT_PROP_NONEXTRA_WATER_SUB_HURT,
	_FightPropLowerName[2250:2284]: FIGHT_PROP_NONEXTRA_WATER_SUB_HURT,
	_FightPropName[2284:2318]:      FIGHT_PROP_NONEXTRA_GRASS_SUB_HURT,
	_FightPropLowerName[2284:2318]: FIGHT_PROP_NONEXTRA_GRASS_SUB_HURT,
	_FightPropName[2318:2351]:      FIGHT_PROP_NONEXTRA_WIND_SUB_HURT,
	_FightPropLowerName[2318:2351]: FIGHT_PROP_NONEXTRA_WIND_SUB_HURT,
	_FightPropName[2351:2384]:      FIGHT_PROP_NONEXTRA_ROCK_SUB_HURT,
	_FightPropLowerName[2351:2384]: FIGHT_PROP_NONEXTRA_ROCK_SUB_HURT,
	_FightPropName[2384:2416]:      FIGHT_PROP_NONEXTRA_ICE_SUB_HURT,
	_FightPropLowerName[2384:2416]: FIGHT_PROP_NONEXTRA_ICE_SUB_HURT,
	_FightPropName[2416:2456]:      FIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIO,
	_FightPropLowerName[2416:2456]: FIGHT_PROP_NONEXTRA_SKILL_CD_MINUS_RATIO,
	_FightPropName[2456:2499]:      FIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIO,
	_FightPropLowerName[2456:2499]: FIGHT_PROP_NONEXTRA_SHIELD_COST_MINUS_RATIO,
	_FightPropName[2499:2536]:      FIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT,
	_FightPropLowerName[2499:2536]: FIGHT_PROP_NONEXTRA_PHYSICAL_ADD_HURT,
}

var _FightPropNames = []string{
	_FightPropName[0:15],
	_FightPropName[15:33],
	_FightPropName[33:46],
	_FightPropName[46:67],
	_FightPropName[67:89],
	_FightPropName[89:106],
	_FightPropName[106:131],
	_FightPropName[131:154],
	_FightPropName[154:172],
	_FightPropName[172:198],
	_FightPropName[198:219],
	_FightPropName[219:243],
	_FightPropName[243:267],
	_FightPropName[267:295],
	_FightPropName[295:314],
	_FightPropName[314:338],
	_FightPropName[338:362],
	_FightPropName[362:390],
	_FightPropName[390:409],
	_FightPropName[409:428],
	_FightPropName[428:447],
	_FightPropName[447:468],
	_FightPropName[468:494],
	_FightPropName[494:522],
	_FightPropName[522:550],
	_FightPropName[550:581],
	_FightPropName[581:612],
	_FightPropName[612:636],
	_FightPropName[636:660],
	_FightPropName[660:685],
	_FightPropName[685:710],
	_FightPropName[710:734],
	_FightPropName[734:758],
	_FightPropName[758:781],
	_FightPropName[781:809],
	_FightPropName[809:833],
	_FightPropName[833:857],
	_FightPropName[857:882],
	_FightPropName[882:907],
	_FightPropName[907:931],
	_FightPropName[931:955],
	_FightPropName[955:978],
	_FightPropName[978:999],
	_FightPropName[999:1023],
	_FightPropName[1023:1047],
	_FightPropName[1047:1071],
	_FightPropName[1071:1094],
	_FightPropName[1094:1119],
	_FightPropName[1119:1144],
	_FightPropName[1144:1168],
	_FightPropName[1168:1194],
	_FightPropName[1194:1220],
	_FightPropName[1220:1247],
	_FightPropName[1247:1274],
	_FightPropName[1274:1300],
	_FightPropName[1300:1325],
	_FightPropName[1325:1351],
	_FightPropName[1351:1382],
	_FightPropName[1382:1416],
	_FightPropName[1416:1442],
	_FightPropName[1442:1468],
	_FightPropName[1468:1495],
	_FightPropName[1495:1522],
	_FightPropName[1522:1548],
	_FightPropName[1548:1573],
	_FightPropName[1573:1599],
	_FightPropName[1599:1616],
	_FightPropName[1616:1633],
	_FightPropName[1633:1654],
	_FightPropName[1654:1676],
	_FightPropName[1676:1696],
	_FightPropName[1696:1722],
	_FightPropName[1722:1749],
	_FightPropName[1749:1777],
	_FightPropName[1777:1810],
	_FightPropName[1810:1843],
	_FightPropName[1843:1880],
	_FightPropName[1880:1915],
	_FightPropName[1915:1952],
	_FightPropName[1952:1985],
	_FightPropName[1985:2018],
	_FightPropName[2018:2052],
	_FightPropName[2052:2086],
	_FightPropName[2086:2119],
	_FightPropName[2119:2152],
	_FightPropName[2152:2184],
	_FightPropName[2184:2217],
	_FightPropName[2217:2250],
	_FightPropName[2250:2284],
	_FightPropName[2284:2318],
	_FightPropName[2318:2351],
	_FightPropName[2351:2384],
	_FightPropName[2384:2416],
	_FightPropName[2416:2456],
	_FightPropName[2456:2499],
	_FightPropName[2499:2536],
}

// FightPropString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FightPropString(s string) (FightProp, error) {
	if val, ok := _FightPropNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FightPropNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FightProp values", s)
}

// FightPropValues returns all values of the enum
func FightPropValues() []FightProp {
	return _FightPropValues
}

// FightPropStrings returns a slice of all String values of the enum
func FightPropStrings() []string {
	strs := make([]string, len(_FightPropNames))
	copy(strs, _FightPropNames)
	return strs
}

// IsAFightProp returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FightProp) IsAFightProp() bool {
	_, ok := _FightPropMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for FightProp
func (i FightProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FightProp
func (i *FightProp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("FightProp should be a string, got %s", data)
	}

	var err error
	*i, err = FightPropString(s)
	return err
}
