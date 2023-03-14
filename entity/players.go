package entity

var TradePlayerMap = map[uint64]uint64{}

type Player struct {
	MaxBid     uint64
	Name       string
}

var Players = map[uint64]Player{
	138956: {MaxBid: 4000, Name: "Giorgio Chiellini"},
	175943: {MaxBid: 4000, Name: "Mertens"},
	220814: {MaxBid: 4000, Name: "Lucas Hernández"},
	228251: {MaxBid: 4000, Name: "Lorenzo Pellegrini"},
	210935: {MaxBid: 4000, Name: "Domenico Berardi"},
	189575: {MaxBid: 4000, Name: "Muniain"},
	233064: {MaxBid: 4000, Name: "Mount"},
	238074: {MaxBid: 4000, Name: "Reece James"},
	181458: {MaxBid: 4000, Name: "Perisic"},
	190286: {MaxBid: 4000, Name: "Canales"},
	188567: {MaxBid: 12000, Name: "Aubameyang"},
	206113: {MaxBid: 12000, Name: "Serge Gnabry"},
	192629: {MaxBid: 12000, Name: "Iago Aspas"},
	224334: {MaxBid: 12000, Name: "Acunha"},
	235790: {MaxBid: 4000, Name: "Kai Havertz"},
	246430: {MaxBid: 4000, Name: "Vlahovic"},
	252371: {MaxBid: 4300, Name: "Jude Bellingham"},
	239053: {MaxBid: 4300, Name: "Federico Valverde"},
	197061: {MaxBid: 4000, Name: "Matip"},
	241084: {MaxBid: 4000, Name: "Luiz Diaz"},
	184344: {MaxBid: 4000, Name: "Leonardo Bonucci"},
	183898: {MaxBid: 4200, Name: "Di Maria"},
	235805: {MaxBid: 4000, Name: "Federico Chiesa"},
	240130: {MaxBid: 6000, Name: "Militao"},
	212616: {MaxBid: 4000, Name: "De Paul"},
	207410: {MaxBid: 3900, Name: "Kovacic"},
	222492: {MaxBid: 3900, Name: "Sané"},
	192984: {MaxBid: 3900, Name: "Casteels"},
	178603: {MaxBid: 3900, Name: "Hummels"},
	183277: {MaxBid: 4000, Name: "Hazar"},
	199451: {MaxBid: 4200, Name: "Ben yedder"},
	177683: {MaxBid: 10000, Name: "Sommer"},
	188943: {MaxBid: 12500, Name: "Kevin Trapp"},
	204963: {MaxBid: 4000, Name: "Carvajal "},
	155862: {MaxBid: 4000, Name: "Sergio Ramos"},
	235212: {MaxBid: 6000, Name: "Hakimi"},
	222665: {MaxBid: 4000, Name: "Ordegard"},
	232756: {MaxBid: 5100, Name: "Tomori"},
	241721: {MaxBid: 4000, Name: "Rafa leão"},
	237383: {MaxBid: 4000, Name: "Bastoni"},
	216393: {MaxBid: 4000, Name: "Tielemans"},
	230142: {MaxBid: 4000, Name: "Oyarzabal "},
	226161: {MaxBid: 4000, Name: "Llorente "},
	230938: {MaxBid: 4000, Name: "Kessié"},
	201535: {MaxBid: 4000, Name: "Varane"},
	202811: {MaxBid: 4000, Name: "Martínez"},
	186345: {MaxBid: 4000, Name: "Trippier"},
	226790: {MaxBid: 4000, Name: "Ndidi"},
	228618: {MaxBid: 4000, Name: "Mendy"},
	143076: {MaxBid: 3900, Name: "Gomez"},
	208670: {MaxBid: 1300, Name: "Ziyech"}, // mudar preço depois se cair
	183711: {MaxBid: 1300, Name: "Henderson"}, // mudar preço depois se cair
	201942: {MaxBid: 1300, Name: "Firmino"}, // mudar preço depois se cair
	239580: {MaxBid: 1300, Name: "Bremer "}, // mudar preço depois se cair
	193082: {MaxBid: 1300, Name: "Cuadrado "}, // mudar preço depois se cair
	230869: {MaxBid: 1300, Name: "Simón"}, // mudar preço depois se cair
	146536: {MaxBid: 1300, Name: "Jesus Navas"}, // mudar preço depois se cair
	241464: {MaxBid: 1300, Name: "Torres"}, // mudar preço depois se cair
}
