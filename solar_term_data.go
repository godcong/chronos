package chronos

var solarTermSanHous = []string{
	0:  "东风解冻、蛰虫始振、鱼上冰。",
	1:  "獭祭鱼、鸿雁来、草木萌动。",
	2:  "桃始花、仓庚鸣、鹰化为鸠。",
	3:  "玄鸟至、雷乃发声、始电。",
	4:  "桐始华、鼠化为鴽、虹始见。",
	5:  "萍始生、鸣鸠拂其羽、戴胜降于桑。",
	6:  "蝼蝈鸣、蚯蚓出、王瓜生。",
	7:  "苦菜秀、靡草死、小暑至。",
	8:  "螳螂生、鵙始鸣、反舌无声。",
	9:  "鹿角解、蜩始鸣、半夏生。",
	10: "温风至、蟋蟀居辟、鹰乃学习。",
	11: "腐草化为萤、土润溽暑、大雨时行。",
	12: "凉风至、白露降、寒蝉鸣。",
	13: "鹰乃祭鸟、天地始肃、禾乃登。",
	14: "鸿雁来、玄鸟归、群鸟养羞。",
	15: "雷始收声、蛰虫培户、水始涸。",
	16: "鸿雁来宾、雀攻大水为蛤、菊有黄花。",
	17: "豺乃祭兽、草木黄落、蛰虫咸俯。",
	18: "水始冰、地始冻、雉入大水为蜃。",
	19: "虹藏不见、天气上腾、闭塞而成冬。",
	20: "鴠鸟不鸣、虎始交、荔挺生。",
	21: "蚯蚓结、麋角解、水泉动。",
	22: "雁北向、鹊始巢、雉始雊。",
	23: "鸡始乳、鸷鸟厉疾、水泽腹坚。",
}
var solarTermExplanations = []string{
	0:  "斗指东北。太阳黄经为315度。是二十四个节气的头一个节气。其含意是开始进入春天，“阳和起蛰，品物皆春”，过了立春，万物复苏生机勃勃，一年四季从此开始了。",
	1:  "斗指壬。太阳黄经为330°。这时春风遍吹，冰雪融化，空气湿润，雨水增多，所以叫雨水。人们常说：“立春天渐暖，雨水送肥忙”。",
	2:  "斗指丁。太阳黄经为345°。这个节气表示“立春”以后天气转暖，春雷开始震响，蛰伏在泥土里的各种冬眠动物将苏醒过来开始活动起来，所以叫惊蛰。这个时期过冬的虫排卵也要开始孵化。我国部分地区过入了春耕季节。谚语云：“惊蛰过，暖和和，蛤蟆老角唱山歌。”“惊蛰一犁土，春分地气通。”“惊蛰没到雷先鸣，大雨似蛟龙。”",
	3:  "斗指壬。太阳黄经为0°。春分日太阳在赤道上方。这是春季90天的中分点，这一天南北两半球昼夜相等，所以叫春分。这天以后太阳直射位置便向北移，北半球昼长夜短。所以春分是北半球春季开始。我国大部分地区越冬作物进入春季生长阶段。各地农谚有：“春分在前，斗米斗钱”（广东）、“春分甲子雨绵绵，夏分甲子火烧天”（四川）、“春分有雨家家忙，先种瓜豆后插秧”（湖北）、“春分种菜， 大暑摘瓜”（湖南）、“春分种麻种豆，秋分种麦种蒜”（安徽）。",
	4:  "斗指丁。太阳黄经为15°。此时气候清爽温暖，草木始发新枝芽，万物开始生长，农民忙于春耕春种。从前，在清明节这一天，有些人家都在门口插上杨柳条，还到郊外踏青，祭扫坟墓，这是古老的习俗。",
	5:  "斗指癸。太阳黄经为30°。就是雨水生五谷的意思，由于雨水滋润大地五谷得以生长，所以，谷雨就是“雨生百谷”。谚云“谷雨前后，种瓜种豆”。",
	6:  "斗指东南。太阳黄经为45°。是夏季的开始，从此进入夏天，万物旺盛大。习惯上把立夏当作是气温显著升高，炎暑将临，雷雨增多，农作物进入旺季生长的一个最重要节气。",
	7:  "斗指甲。太阳黄经为60°。从小满开始，大麦、冬小麦等夏收作物，已经结果、籽粒饱满，但尚未成熟，所以叫小满。",
	8:  "北斗指向已。太阳黄经为75°。这时最适合播种有芒的谷类作物，如晚谷、黍、稷等。如过了这个时候再种有芒和作物就不好成熟了。同时，“芒”指有芒作物如小麦、大麦等，“种”指种子。芒种即表明小麦等有芒作物成熟。芒种前后，我国中部的长江中、下游地区，雨量增多，气温升高，进入连绵阴雨的梅雨季节，空气非常潮湿，天气异常闷热，各种器具和衣物容易发霉，所以在我国长江中、下游地 区也叫“霉雨”。",
	9:  "北斗指向乙。太阳黄经为90°。太阳在黄经90°“夏至点”时，阳光几乎直射北回归线上空，北半球正午太阳最高。这一天是北半球白昼最长、黑夜最短的一天，从这一天起，进入炎热季节，天地万物在此时生长最旺盛。所心以古时候又把这一天叫做日北至，意思是太阳运生到最北的一日。过了夏至，太阳逐渐向南移动，北半球白昼一天比一天缩短，黑夜一天比一天加长。",
	10: "斗指辛。太阳黄经为105°。天气已经很热，但不到是热的时候，所以叫小暑。此时，已是初伏前后。",
	11: "斗指丙。太阳黄经为120°。大暑是一年中最热的节气，正值勤二伏前后，长江流域的许多地方，经常出现40℃高温天气。要作好防暑降温工作。这个节气雨水多，在“小暑、大暑，淹死老鼠”的谚语，要注意防汛防涝。",
	12: "北斗指向西南。太阳黄经为135°。从这一天起秋天开始，秋高气爽，月明风清。此后，气温由最热逐渐下降。",
	13: "斗指戊。太阳黄经为150°。这时夏季火热已经到头了。暑气就要散了。它是温度下降的一个转折点。是气候变凉的象征，表示暑天终止。",
	14: "斗指癸。太阳黄经为165°。天气转凉，地面水汽结露最多。",
	15: "斗指已。太阳黄经为180°。秋分这一天同春人一样，阳光几乎直射赤道，昼夜几乎相等。从这一天起，阳光直射位置继续由赤道向南半球推移，北半球开始昼短夜长。依我国旧历的秋季论，这一天刚好是秋季九十天的一半，因而称秋分。但在天文学上规定，北半球的秋天是从秋分开始的。",
	16: "斗指甲。太阳黄经为195°。白露后，天气转凉，开始出现露水，到了寒露，则露水日多，且气温更低了。所以，有人说，寒是露之气，先白而后寒，是气候将逐渐转冷的意思。而水气则凝成白色露珠。",
	17: "太阳黄经为210°。天气已冷，开始有霜冻了，所以叫霜降。",
	18: "太阳黄经为225°。习惯上，我国人民把这一天当作冬季的开始。冬，作为终了之意，是指一年的田间操作结束了，作物收割之后要收藏起来的意思。立冬一过，我国黄河中、下游地区即将结冰，我国各地农民都将陆续地转入农田水利基本建设和其他农事活动中。",
	19: "太阳黄经为240°。气温下降，开始降雪，但还不到大雪纷飞的时节，所以叫小雪。小雪前后，黄河流域开始降雪（南方降雪还要晚两个节气）；而北方，已进入封冻季节。",
	20: "太阳黄经为255°。大雪前后，黄河流域一带渐有积雪；而北方，已是“千里冰封，万里雪飘荡”的严冬了。",
	21: "太阳黄经为270°。冬至这一天，阳光几乎直射南回归线，我们北半球白昼最短，黑夜最长，开始进入数九寒天。天文学上规定这一天是北半球冬季的开始。而冬至以后，阳光直射位置逐渐向北移动，北半球的白天就逐渐长了，谚云：吃了冬至面，一天长一线。",
	22: "太阳黄经为285°。小寒以后，开始进入寒冷季节。冷气积久而寒，小寒是天气寒冷但还没有到极点的意思。",
	23: "太阳黄经为300°。大寒就是天气寒冷到了极点的意思。大寒前后是一年中最冷的季节。大寒正值三九刚过，四九之初。谚云：“三九四九冰上走”。",
}

var solarTermTimes = map[int][24]int64{
	//1900: "ffffffff7c831763ffffffff7c96a7eaffffffff7caa5720ffffffff7cbe2fb5ffffffff7cd231a9ffffffff7ce662aaffffffff7cfabc60ffffffff7d0f3ec7ffffffff7d23dd5fffffffff7d389321ffffffff7d4d4fd0ffffffff7d620b77ffffffff7d76b65affffffff7d8b4895ffffffff7d9fb716ffffffff7db3fcfbffffffff7dc81635ffffffff7ddc02b4ffffffff7defc590ffffffff7e0363c6ffffffff7e16e5d6ffffffff7e2a549effffffff7c5c360dffffffff7c6fa0c9",
	1900: [24]int64{
		0xFFFFFFFF7C831763,
		0xFFFFFFFF7C96A7EA,
		0xFFFFFFFF7CAA5720,
		0xFFFFFFFF7CBE2FB5,
		0xFFFFFFFF7CD231A9,
		0xFFFFFFFF7CE662AA,
		0xFFFFFFFF7CFABC60,
		0xFFFFFFFF7D0F3EC7,
		0xFFFFFFFF7D23DD5F,
		0xFFFFFFFF7D389321,
		0xFFFFFFFF7D4D4FD0,
		0xFFFFFFFF7D620B77,
		0xFFFFFFFF7D76B65A,
		0xFFFFFFFF7D8B4895,
		0xFFFFFFFF7D9FB716,
		0xFFFFFFFF7DB3FCFB,
		0xFFFFFFFF7DC81635,
		0xFFFFFFFF7DDC02B4,
		0xFFFFFFFF7DEFC590,
		0xFFFFFFFF7E0363C6,
		0xFFFFFFFF7E16E5D6,
		0xFFFFFFFF7E2A549E,
		0xFFFFFFFF7C5C360D,
		0xFFFFFFFF7C6FA0C9,
	},
}
