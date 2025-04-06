// Code generated by "enumer --json --linecomment --type=CharacterKey keys_gen.go"; DO NOT EDIT.

package good

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	_CharacterKeyName_0      = ""
	_CharacterKeyLowerName_0 = ""
	_CharacterKeyName_1      = "KamisatoAyakaJean"
	_CharacterKeyLowerName_1 = "kamisatoayakajean"
	_CharacterKeyName_2      = "LisaTraveler"
	_CharacterKeyLowerName_2 = "lisatraveler"
	_CharacterKeyName_3      = "BarbaraKaeyaDiluc"
	_CharacterKeyLowerName_3 = "barbarakaeyadiluc"
	_CharacterKeyName_4      = "RazorAmberVentiXianglingBeidouXingqiuXiaoNingguang"
	_CharacterKeyLowerName_4 = "razoramberventixianglingbeidouxingqiuxiaoningguang"
	_CharacterKeyName_5      = "KleeZhongliFischlBennettTartagliaNoelleQiqiChongyunGanyuAlbedoDiona"
	_CharacterKeyLowerName_5 = "kleezhonglifischlbennetttartaglianoelleqiqichongyunganyualbedodiona"
	_CharacterKeyName_6      = "MonaKeqingSucroseXinyanRosariaHuTaoKaedeharaKazuhaYanfeiYoimiyaThomaEulaRaidenShogunSayuSangonomiyaKokomiGorouKujouSaraAratakiIttoYaeMikoShikanoinHeizouYelanKiraraAloyShenheYunJinKukiShinobuKamisatoAyatoColleiDoriTighnariNilouCynoCandaceNahidaLaylaWandererFaruzanYaoyaoAlhaithamDehyaMikaKavehBaizhuLynetteLyneyFreminetWriothesleyNeuvilletteCharlotteFurinaChevreuseNaviaGamingXianyunChioriSigewinneArlecchinoSethosClorindeEmilieKachinaKinichMualaniXilonenChascaOroronMavuikaCitlaliLanYanYumemizukiMizukiIansanVaresa"
	_CharacterKeyLowerName_6 = "monakeqingsucrosexinyanrosariahutaokaedeharakazuhayanfeiyoimiyathomaeularaidenshogunsayusangonomiyakokomigoroukujousaraaratakiittoyaemikoshikanoinheizouyelankiraraaloyshenheyunjinkukishinobukamisatoayatocolleidoritighnariniloucynocandacenahidalaylawandererfaruzanyaoyaoalhaithamdehyamikakavehbaizhulynettelyneyfreminetwriothesleyneuvillettecharlottefurinachevreusenaviagamingxianyunchiorisigewinnearlecchinosethosclorindeemiliekachinakinichmualanixilonenchascaororonmavuikacitlalilanyanyumemizukimizukiiansanvaresa"
)

var (
	_CharacterKeyIndex_0 = [...]uint8{0, 0}
	_CharacterKeyIndex_1 = [...]uint8{0, 13, 17}
	_CharacterKeyIndex_2 = [...]uint8{0, 4, 12}
	_CharacterKeyIndex_3 = [...]uint8{0, 7, 12, 17}
	_CharacterKeyIndex_4 = [...]uint8{0, 5, 10, 15, 24, 30, 37, 41, 50}
	_CharacterKeyIndex_5 = [...]uint8{0, 4, 11, 17, 24, 33, 39, 43, 51, 56, 62, 67}
	_CharacterKeyIndex_6 = [...]uint16{0, 4, 10, 17, 23, 30, 35, 50, 56, 63, 68, 72, 84, 88, 105, 110, 119, 130, 137, 152, 157, 163, 167, 173, 179, 190, 203, 209, 213, 221, 226, 230, 237, 243, 248, 256, 263, 269, 278, 283, 287, 292, 298, 305, 310, 318, 329, 340, 349, 355, 364, 369, 375, 382, 388, 397, 407, 413, 421, 427, 434, 440, 447, 454, 460, 466, 473, 480, 486, 502, 508, 514}
)

func (i CharacterKey) String() string {
	switch {
	case i == 0:
		return _CharacterKeyName_0
	case 10000002 <= i && i <= 10000003:
		i -= 10000002
		return _CharacterKeyName_1[_CharacterKeyIndex_1[i]:_CharacterKeyIndex_1[i+1]]
	case 10000006 <= i && i <= 10000007:
		i -= 10000006
		return _CharacterKeyName_2[_CharacterKeyIndex_2[i]:_CharacterKeyIndex_2[i+1]]
	case 10000014 <= i && i <= 10000016:
		i -= 10000014
		return _CharacterKeyName_3[_CharacterKeyIndex_3[i]:_CharacterKeyIndex_3[i+1]]
	case 10000020 <= i && i <= 10000027:
		i -= 10000020
		return _CharacterKeyName_4[_CharacterKeyIndex_4[i]:_CharacterKeyIndex_4[i+1]]
	case 10000029 <= i && i <= 10000039:
		i -= 10000029
		return _CharacterKeyName_5[_CharacterKeyIndex_5[i]:_CharacterKeyIndex_5[i+1]]
	case 10000041 <= i && i <= 10000111:
		i -= 10000041
		return _CharacterKeyName_6[_CharacterKeyIndex_6[i]:_CharacterKeyIndex_6[i+1]]
	default:
		return fmt.Sprintf("CharacterKey(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _CharacterKeyNoOp() {
	var x [1]struct{}
	_ = x[UnknownCharacterKey-(0)]
	_ = x[KamisatoAyaka-(10000002)]
	_ = x[Jean-(10000003)]
	_ = x[Lisa-(10000006)]
	_ = x[Traveler-(10000007)]
	_ = x[Barbara-(10000014)]
	_ = x[Kaeya-(10000015)]
	_ = x[Diluc-(10000016)]
	_ = x[Razor-(10000020)]
	_ = x[Amber-(10000021)]
	_ = x[Venti-(10000022)]
	_ = x[Xiangling-(10000023)]
	_ = x[Beidou-(10000024)]
	_ = x[Xingqiu-(10000025)]
	_ = x[Xiao-(10000026)]
	_ = x[Ningguang-(10000027)]
	_ = x[Klee-(10000029)]
	_ = x[Zhongli-(10000030)]
	_ = x[Fischl-(10000031)]
	_ = x[Bennett-(10000032)]
	_ = x[Tartaglia-(10000033)]
	_ = x[Noelle-(10000034)]
	_ = x[Qiqi-(10000035)]
	_ = x[Chongyun-(10000036)]
	_ = x[Ganyu-(10000037)]
	_ = x[Albedo-(10000038)]
	_ = x[Diona-(10000039)]
	_ = x[Mona-(10000041)]
	_ = x[Keqing-(10000042)]
	_ = x[Sucrose-(10000043)]
	_ = x[Xinyan-(10000044)]
	_ = x[Rosaria-(10000045)]
	_ = x[HuTao-(10000046)]
	_ = x[KaedeharaKazuha-(10000047)]
	_ = x[Yanfei-(10000048)]
	_ = x[Yoimiya-(10000049)]
	_ = x[Thoma-(10000050)]
	_ = x[Eula-(10000051)]
	_ = x[RaidenShogun-(10000052)]
	_ = x[Sayu-(10000053)]
	_ = x[SangonomiyaKokomi-(10000054)]
	_ = x[Gorou-(10000055)]
	_ = x[KujouSara-(10000056)]
	_ = x[AratakiItto-(10000057)]
	_ = x[YaeMiko-(10000058)]
	_ = x[ShikanoinHeizou-(10000059)]
	_ = x[Yelan-(10000060)]
	_ = x[Kirara-(10000061)]
	_ = x[Aloy-(10000062)]
	_ = x[Shenhe-(10000063)]
	_ = x[YunJin-(10000064)]
	_ = x[KukiShinobu-(10000065)]
	_ = x[KamisatoAyato-(10000066)]
	_ = x[Collei-(10000067)]
	_ = x[Dori-(10000068)]
	_ = x[Tighnari-(10000069)]
	_ = x[Nilou-(10000070)]
	_ = x[Cyno-(10000071)]
	_ = x[Candace-(10000072)]
	_ = x[Nahida-(10000073)]
	_ = x[Layla-(10000074)]
	_ = x[Wanderer-(10000075)]
	_ = x[Faruzan-(10000076)]
	_ = x[Yaoyao-(10000077)]
	_ = x[Alhaitham-(10000078)]
	_ = x[Dehya-(10000079)]
	_ = x[Mika-(10000080)]
	_ = x[Kaveh-(10000081)]
	_ = x[Baizhu-(10000082)]
	_ = x[Lynette-(10000083)]
	_ = x[Lyney-(10000084)]
	_ = x[Freminet-(10000085)]
	_ = x[Wriothesley-(10000086)]
	_ = x[Neuvillette-(10000087)]
	_ = x[Charlotte-(10000088)]
	_ = x[Furina-(10000089)]
	_ = x[Chevreuse-(10000090)]
	_ = x[Navia-(10000091)]
	_ = x[Gaming-(10000092)]
	_ = x[Xianyun-(10000093)]
	_ = x[Chiori-(10000094)]
	_ = x[Sigewinne-(10000095)]
	_ = x[Arlecchino-(10000096)]
	_ = x[Sethos-(10000097)]
	_ = x[Clorinde-(10000098)]
	_ = x[Emilie-(10000099)]
	_ = x[Kachina-(10000100)]
	_ = x[Kinich-(10000101)]
	_ = x[Mualani-(10000102)]
	_ = x[Xilonen-(10000103)]
	_ = x[Chasca-(10000104)]
	_ = x[Ororon-(10000105)]
	_ = x[Mavuika-(10000106)]
	_ = x[Citlali-(10000107)]
	_ = x[LanYan-(10000108)]
	_ = x[YumemizukiMizuki-(10000109)]
	_ = x[Iansan-(10000110)]
	_ = x[Varesa-(10000111)]
}

var _CharacterKeyValues = []CharacterKey{UnknownCharacterKey, KamisatoAyaka, Jean, Lisa, Traveler, Barbara, Kaeya, Diluc, Razor, Amber, Venti, Xiangling, Beidou, Xingqiu, Xiao, Ningguang, Klee, Zhongli, Fischl, Bennett, Tartaglia, Noelle, Qiqi, Chongyun, Ganyu, Albedo, Diona, Mona, Keqing, Sucrose, Xinyan, Rosaria, HuTao, KaedeharaKazuha, Yanfei, Yoimiya, Thoma, Eula, RaidenShogun, Sayu, SangonomiyaKokomi, Gorou, KujouSara, AratakiItto, YaeMiko, ShikanoinHeizou, Yelan, Kirara, Aloy, Shenhe, YunJin, KukiShinobu, KamisatoAyato, Collei, Dori, Tighnari, Nilou, Cyno, Candace, Nahida, Layla, Wanderer, Faruzan, Yaoyao, Alhaitham, Dehya, Mika, Kaveh, Baizhu, Lynette, Lyney, Freminet, Wriothesley, Neuvillette, Charlotte, Furina, Chevreuse, Navia, Gaming, Xianyun, Chiori, Sigewinne, Arlecchino, Sethos, Clorinde, Emilie, Kachina, Kinich, Mualani, Xilonen, Chasca, Ororon, Mavuika, Citlali, LanYan, YumemizukiMizuki, Iansan, Varesa}

var _CharacterKeyNameToValueMap = map[string]CharacterKey{
	_CharacterKeyName_0[0:0]:          UnknownCharacterKey,
	_CharacterKeyLowerName_0[0:0]:     UnknownCharacterKey,
	_CharacterKeyName_1[0:13]:         KamisatoAyaka,
	_CharacterKeyLowerName_1[0:13]:    KamisatoAyaka,
	_CharacterKeyName_1[13:17]:        Jean,
	_CharacterKeyLowerName_1[13:17]:   Jean,
	_CharacterKeyName_2[0:4]:          Lisa,
	_CharacterKeyLowerName_2[0:4]:     Lisa,
	_CharacterKeyName_2[4:12]:         Traveler,
	_CharacterKeyLowerName_2[4:12]:    Traveler,
	_CharacterKeyName_3[0:7]:          Barbara,
	_CharacterKeyLowerName_3[0:7]:     Barbara,
	_CharacterKeyName_3[7:12]:         Kaeya,
	_CharacterKeyLowerName_3[7:12]:    Kaeya,
	_CharacterKeyName_3[12:17]:        Diluc,
	_CharacterKeyLowerName_3[12:17]:   Diluc,
	_CharacterKeyName_4[0:5]:          Razor,
	_CharacterKeyLowerName_4[0:5]:     Razor,
	_CharacterKeyName_4[5:10]:         Amber,
	_CharacterKeyLowerName_4[5:10]:    Amber,
	_CharacterKeyName_4[10:15]:        Venti,
	_CharacterKeyLowerName_4[10:15]:   Venti,
	_CharacterKeyName_4[15:24]:        Xiangling,
	_CharacterKeyLowerName_4[15:24]:   Xiangling,
	_CharacterKeyName_4[24:30]:        Beidou,
	_CharacterKeyLowerName_4[24:30]:   Beidou,
	_CharacterKeyName_4[30:37]:        Xingqiu,
	_CharacterKeyLowerName_4[30:37]:   Xingqiu,
	_CharacterKeyName_4[37:41]:        Xiao,
	_CharacterKeyLowerName_4[37:41]:   Xiao,
	_CharacterKeyName_4[41:50]:        Ningguang,
	_CharacterKeyLowerName_4[41:50]:   Ningguang,
	_CharacterKeyName_5[0:4]:          Klee,
	_CharacterKeyLowerName_5[0:4]:     Klee,
	_CharacterKeyName_5[4:11]:         Zhongli,
	_CharacterKeyLowerName_5[4:11]:    Zhongli,
	_CharacterKeyName_5[11:17]:        Fischl,
	_CharacterKeyLowerName_5[11:17]:   Fischl,
	_CharacterKeyName_5[17:24]:        Bennett,
	_CharacterKeyLowerName_5[17:24]:   Bennett,
	_CharacterKeyName_5[24:33]:        Tartaglia,
	_CharacterKeyLowerName_5[24:33]:   Tartaglia,
	_CharacterKeyName_5[33:39]:        Noelle,
	_CharacterKeyLowerName_5[33:39]:   Noelle,
	_CharacterKeyName_5[39:43]:        Qiqi,
	_CharacterKeyLowerName_5[39:43]:   Qiqi,
	_CharacterKeyName_5[43:51]:        Chongyun,
	_CharacterKeyLowerName_5[43:51]:   Chongyun,
	_CharacterKeyName_5[51:56]:        Ganyu,
	_CharacterKeyLowerName_5[51:56]:   Ganyu,
	_CharacterKeyName_5[56:62]:        Albedo,
	_CharacterKeyLowerName_5[56:62]:   Albedo,
	_CharacterKeyName_5[62:67]:        Diona,
	_CharacterKeyLowerName_5[62:67]:   Diona,
	_CharacterKeyName_6[0:4]:          Mona,
	_CharacterKeyLowerName_6[0:4]:     Mona,
	_CharacterKeyName_6[4:10]:         Keqing,
	_CharacterKeyLowerName_6[4:10]:    Keqing,
	_CharacterKeyName_6[10:17]:        Sucrose,
	_CharacterKeyLowerName_6[10:17]:   Sucrose,
	_CharacterKeyName_6[17:23]:        Xinyan,
	_CharacterKeyLowerName_6[17:23]:   Xinyan,
	_CharacterKeyName_6[23:30]:        Rosaria,
	_CharacterKeyLowerName_6[23:30]:   Rosaria,
	_CharacterKeyName_6[30:35]:        HuTao,
	_CharacterKeyLowerName_6[30:35]:   HuTao,
	_CharacterKeyName_6[35:50]:        KaedeharaKazuha,
	_CharacterKeyLowerName_6[35:50]:   KaedeharaKazuha,
	_CharacterKeyName_6[50:56]:        Yanfei,
	_CharacterKeyLowerName_6[50:56]:   Yanfei,
	_CharacterKeyName_6[56:63]:        Yoimiya,
	_CharacterKeyLowerName_6[56:63]:   Yoimiya,
	_CharacterKeyName_6[63:68]:        Thoma,
	_CharacterKeyLowerName_6[63:68]:   Thoma,
	_CharacterKeyName_6[68:72]:        Eula,
	_CharacterKeyLowerName_6[68:72]:   Eula,
	_CharacterKeyName_6[72:84]:        RaidenShogun,
	_CharacterKeyLowerName_6[72:84]:   RaidenShogun,
	_CharacterKeyName_6[84:88]:        Sayu,
	_CharacterKeyLowerName_6[84:88]:   Sayu,
	_CharacterKeyName_6[88:105]:       SangonomiyaKokomi,
	_CharacterKeyLowerName_6[88:105]:  SangonomiyaKokomi,
	_CharacterKeyName_6[105:110]:      Gorou,
	_CharacterKeyLowerName_6[105:110]: Gorou,
	_CharacterKeyName_6[110:119]:      KujouSara,
	_CharacterKeyLowerName_6[110:119]: KujouSara,
	_CharacterKeyName_6[119:130]:      AratakiItto,
	_CharacterKeyLowerName_6[119:130]: AratakiItto,
	_CharacterKeyName_6[130:137]:      YaeMiko,
	_CharacterKeyLowerName_6[130:137]: YaeMiko,
	_CharacterKeyName_6[137:152]:      ShikanoinHeizou,
	_CharacterKeyLowerName_6[137:152]: ShikanoinHeizou,
	_CharacterKeyName_6[152:157]:      Yelan,
	_CharacterKeyLowerName_6[152:157]: Yelan,
	_CharacterKeyName_6[157:163]:      Kirara,
	_CharacterKeyLowerName_6[157:163]: Kirara,
	_CharacterKeyName_6[163:167]:      Aloy,
	_CharacterKeyLowerName_6[163:167]: Aloy,
	_CharacterKeyName_6[167:173]:      Shenhe,
	_CharacterKeyLowerName_6[167:173]: Shenhe,
	_CharacterKeyName_6[173:179]:      YunJin,
	_CharacterKeyLowerName_6[173:179]: YunJin,
	_CharacterKeyName_6[179:190]:      KukiShinobu,
	_CharacterKeyLowerName_6[179:190]: KukiShinobu,
	_CharacterKeyName_6[190:203]:      KamisatoAyato,
	_CharacterKeyLowerName_6[190:203]: KamisatoAyato,
	_CharacterKeyName_6[203:209]:      Collei,
	_CharacterKeyLowerName_6[203:209]: Collei,
	_CharacterKeyName_6[209:213]:      Dori,
	_CharacterKeyLowerName_6[209:213]: Dori,
	_CharacterKeyName_6[213:221]:      Tighnari,
	_CharacterKeyLowerName_6[213:221]: Tighnari,
	_CharacterKeyName_6[221:226]:      Nilou,
	_CharacterKeyLowerName_6[221:226]: Nilou,
	_CharacterKeyName_6[226:230]:      Cyno,
	_CharacterKeyLowerName_6[226:230]: Cyno,
	_CharacterKeyName_6[230:237]:      Candace,
	_CharacterKeyLowerName_6[230:237]: Candace,
	_CharacterKeyName_6[237:243]:      Nahida,
	_CharacterKeyLowerName_6[237:243]: Nahida,
	_CharacterKeyName_6[243:248]:      Layla,
	_CharacterKeyLowerName_6[243:248]: Layla,
	_CharacterKeyName_6[248:256]:      Wanderer,
	_CharacterKeyLowerName_6[248:256]: Wanderer,
	_CharacterKeyName_6[256:263]:      Faruzan,
	_CharacterKeyLowerName_6[256:263]: Faruzan,
	_CharacterKeyName_6[263:269]:      Yaoyao,
	_CharacterKeyLowerName_6[263:269]: Yaoyao,
	_CharacterKeyName_6[269:278]:      Alhaitham,
	_CharacterKeyLowerName_6[269:278]: Alhaitham,
	_CharacterKeyName_6[278:283]:      Dehya,
	_CharacterKeyLowerName_6[278:283]: Dehya,
	_CharacterKeyName_6[283:287]:      Mika,
	_CharacterKeyLowerName_6[283:287]: Mika,
	_CharacterKeyName_6[287:292]:      Kaveh,
	_CharacterKeyLowerName_6[287:292]: Kaveh,
	_CharacterKeyName_6[292:298]:      Baizhu,
	_CharacterKeyLowerName_6[292:298]: Baizhu,
	_CharacterKeyName_6[298:305]:      Lynette,
	_CharacterKeyLowerName_6[298:305]: Lynette,
	_CharacterKeyName_6[305:310]:      Lyney,
	_CharacterKeyLowerName_6[305:310]: Lyney,
	_CharacterKeyName_6[310:318]:      Freminet,
	_CharacterKeyLowerName_6[310:318]: Freminet,
	_CharacterKeyName_6[318:329]:      Wriothesley,
	_CharacterKeyLowerName_6[318:329]: Wriothesley,
	_CharacterKeyName_6[329:340]:      Neuvillette,
	_CharacterKeyLowerName_6[329:340]: Neuvillette,
	_CharacterKeyName_6[340:349]:      Charlotte,
	_CharacterKeyLowerName_6[340:349]: Charlotte,
	_CharacterKeyName_6[349:355]:      Furina,
	_CharacterKeyLowerName_6[349:355]: Furina,
	_CharacterKeyName_6[355:364]:      Chevreuse,
	_CharacterKeyLowerName_6[355:364]: Chevreuse,
	_CharacterKeyName_6[364:369]:      Navia,
	_CharacterKeyLowerName_6[364:369]: Navia,
	_CharacterKeyName_6[369:375]:      Gaming,
	_CharacterKeyLowerName_6[369:375]: Gaming,
	_CharacterKeyName_6[375:382]:      Xianyun,
	_CharacterKeyLowerName_6[375:382]: Xianyun,
	_CharacterKeyName_6[382:388]:      Chiori,
	_CharacterKeyLowerName_6[382:388]: Chiori,
	_CharacterKeyName_6[388:397]:      Sigewinne,
	_CharacterKeyLowerName_6[388:397]: Sigewinne,
	_CharacterKeyName_6[397:407]:      Arlecchino,
	_CharacterKeyLowerName_6[397:407]: Arlecchino,
	_CharacterKeyName_6[407:413]:      Sethos,
	_CharacterKeyLowerName_6[407:413]: Sethos,
	_CharacterKeyName_6[413:421]:      Clorinde,
	_CharacterKeyLowerName_6[413:421]: Clorinde,
	_CharacterKeyName_6[421:427]:      Emilie,
	_CharacterKeyLowerName_6[421:427]: Emilie,
	_CharacterKeyName_6[427:434]:      Kachina,
	_CharacterKeyLowerName_6[427:434]: Kachina,
	_CharacterKeyName_6[434:440]:      Kinich,
	_CharacterKeyLowerName_6[434:440]: Kinich,
	_CharacterKeyName_6[440:447]:      Mualani,
	_CharacterKeyLowerName_6[440:447]: Mualani,
	_CharacterKeyName_6[447:454]:      Xilonen,
	_CharacterKeyLowerName_6[447:454]: Xilonen,
	_CharacterKeyName_6[454:460]:      Chasca,
	_CharacterKeyLowerName_6[454:460]: Chasca,
	_CharacterKeyName_6[460:466]:      Ororon,
	_CharacterKeyLowerName_6[460:466]: Ororon,
	_CharacterKeyName_6[466:473]:      Mavuika,
	_CharacterKeyLowerName_6[466:473]: Mavuika,
	_CharacterKeyName_6[473:480]:      Citlali,
	_CharacterKeyLowerName_6[473:480]: Citlali,
	_CharacterKeyName_6[480:486]:      LanYan,
	_CharacterKeyLowerName_6[480:486]: LanYan,
	_CharacterKeyName_6[486:502]:      YumemizukiMizuki,
	_CharacterKeyLowerName_6[486:502]: YumemizukiMizuki,
	_CharacterKeyName_6[502:508]:      Iansan,
	_CharacterKeyLowerName_6[502:508]: Iansan,
	_CharacterKeyName_6[508:514]:      Varesa,
	_CharacterKeyLowerName_6[508:514]: Varesa,
}

var _CharacterKeyNames = []string{
	_CharacterKeyName_0[0:0],
	_CharacterKeyName_1[0:13],
	_CharacterKeyName_1[13:17],
	_CharacterKeyName_2[0:4],
	_CharacterKeyName_2[4:12],
	_CharacterKeyName_3[0:7],
	_CharacterKeyName_3[7:12],
	_CharacterKeyName_3[12:17],
	_CharacterKeyName_4[0:5],
	_CharacterKeyName_4[5:10],
	_CharacterKeyName_4[10:15],
	_CharacterKeyName_4[15:24],
	_CharacterKeyName_4[24:30],
	_CharacterKeyName_4[30:37],
	_CharacterKeyName_4[37:41],
	_CharacterKeyName_4[41:50],
	_CharacterKeyName_5[0:4],
	_CharacterKeyName_5[4:11],
	_CharacterKeyName_5[11:17],
	_CharacterKeyName_5[17:24],
	_CharacterKeyName_5[24:33],
	_CharacterKeyName_5[33:39],
	_CharacterKeyName_5[39:43],
	_CharacterKeyName_5[43:51],
	_CharacterKeyName_5[51:56],
	_CharacterKeyName_5[56:62],
	_CharacterKeyName_5[62:67],
	_CharacterKeyName_6[0:4],
	_CharacterKeyName_6[4:10],
	_CharacterKeyName_6[10:17],
	_CharacterKeyName_6[17:23],
	_CharacterKeyName_6[23:30],
	_CharacterKeyName_6[30:35],
	_CharacterKeyName_6[35:50],
	_CharacterKeyName_6[50:56],
	_CharacterKeyName_6[56:63],
	_CharacterKeyName_6[63:68],
	_CharacterKeyName_6[68:72],
	_CharacterKeyName_6[72:84],
	_CharacterKeyName_6[84:88],
	_CharacterKeyName_6[88:105],
	_CharacterKeyName_6[105:110],
	_CharacterKeyName_6[110:119],
	_CharacterKeyName_6[119:130],
	_CharacterKeyName_6[130:137],
	_CharacterKeyName_6[137:152],
	_CharacterKeyName_6[152:157],
	_CharacterKeyName_6[157:163],
	_CharacterKeyName_6[163:167],
	_CharacterKeyName_6[167:173],
	_CharacterKeyName_6[173:179],
	_CharacterKeyName_6[179:190],
	_CharacterKeyName_6[190:203],
	_CharacterKeyName_6[203:209],
	_CharacterKeyName_6[209:213],
	_CharacterKeyName_6[213:221],
	_CharacterKeyName_6[221:226],
	_CharacterKeyName_6[226:230],
	_CharacterKeyName_6[230:237],
	_CharacterKeyName_6[237:243],
	_CharacterKeyName_6[243:248],
	_CharacterKeyName_6[248:256],
	_CharacterKeyName_6[256:263],
	_CharacterKeyName_6[263:269],
	_CharacterKeyName_6[269:278],
	_CharacterKeyName_6[278:283],
	_CharacterKeyName_6[283:287],
	_CharacterKeyName_6[287:292],
	_CharacterKeyName_6[292:298],
	_CharacterKeyName_6[298:305],
	_CharacterKeyName_6[305:310],
	_CharacterKeyName_6[310:318],
	_CharacterKeyName_6[318:329],
	_CharacterKeyName_6[329:340],
	_CharacterKeyName_6[340:349],
	_CharacterKeyName_6[349:355],
	_CharacterKeyName_6[355:364],
	_CharacterKeyName_6[364:369],
	_CharacterKeyName_6[369:375],
	_CharacterKeyName_6[375:382],
	_CharacterKeyName_6[382:388],
	_CharacterKeyName_6[388:397],
	_CharacterKeyName_6[397:407],
	_CharacterKeyName_6[407:413],
	_CharacterKeyName_6[413:421],
	_CharacterKeyName_6[421:427],
	_CharacterKeyName_6[427:434],
	_CharacterKeyName_6[434:440],
	_CharacterKeyName_6[440:447],
	_CharacterKeyName_6[447:454],
	_CharacterKeyName_6[454:460],
	_CharacterKeyName_6[460:466],
	_CharacterKeyName_6[466:473],
	_CharacterKeyName_6[473:480],
	_CharacterKeyName_6[480:486],
	_CharacterKeyName_6[486:502],
	_CharacterKeyName_6[502:508],
	_CharacterKeyName_6[508:514],
}

// CharacterKeyString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CharacterKeyString(s string) (CharacterKey, error) {
	if val, ok := _CharacterKeyNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _CharacterKeyNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CharacterKey values", s)
}

// CharacterKeyValues returns all values of the enum
func CharacterKeyValues() []CharacterKey {
	return _CharacterKeyValues
}

// CharacterKeyStrings returns a slice of all String values of the enum
func CharacterKeyStrings() []string {
	strs := make([]string, len(_CharacterKeyNames))
	copy(strs, _CharacterKeyNames)
	return strs
}

// IsACharacterKey returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CharacterKey) IsACharacterKey() bool {
	for _, v := range _CharacterKeyValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CharacterKey
func (i CharacterKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CharacterKey
func (i *CharacterKey) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CharacterKey should be a string, got %s", data)
	}

	var err error
	*i, err = CharacterKeyString(s)
	return err
}
