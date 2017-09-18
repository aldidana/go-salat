package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"

	"github.com/urfave/cli"
)

var j = []byte(`{"Ambarawa":"1","Ambon":"2","Amlapura":"3","Amuntai":"4","Argamakmur":"5","Atambua":"6","Babo":"7","Bagan Siapiapi":"8","Bajawa":"9","Balige":"10","Balik Papan":"11","Banda Aceh":"12","Bandarlampung":"13","Bandung":"14","Bangkalan":"15","Bangkinang":"16","Bangko":"17","Bangli":"18","Banjar":"19","Banjar Baru":"20","Banjarmasin":"21","Banjarnegara":"22","Bantaeng":"23","Banten":"24","Bantul":"25","Banyuwangi":"26","Barabai":"27","Barito":"28","Barru":"29","Batam":"30","Batang":"31","Batu":"32","Baturaja":"33","Batusangkar":"34","Baubau":"35","Bekasi":"36","Bengkalis":"37","Bengkulu":"38","Benteng":"39","Biak":"40","Bima":"41","Binjai":"42","Bireuen":"43","Bitung":"44","Blitar":"45","Blora":"46","Bogor":"47","Bojonegoro":"48","Bondowoso":"49","Bontang":"50","Boyolali":"51","Brebes":"52","Bukit Tinggi":"53","Bulukumba":"54","Buntok":"55","Cepu":"56","Ciamis":"57","Cianjur":"58","Cibinong":"59","Cilacap":"60","Cilegon":"61","Cimahi":"62","Cirebon":"63","Curup":"64","Demak":"65","Denpasar":"66","Depok":"67","Dili":"68","Dompu":"69","Donggala":"70","Dumai":"71","Ende":"72","Enggano":"73","Enrekang":"74","Fakfak":"75","Garut":"76","Gianyar":"77","Gombong":"78","Gorontalo":"79","Gresik":"80","Gunung Sitoli":"81","Indramayu":"82","Jakarta":"83","Jambi":"84","Jayapura":"85","Jember":"86","Jeneponto":"87","Jepara":"88","Jombang":"89","Kabanjahe":"90","Kalabahi":"91","Kalianda":"92","Kandangan":"93","Karanganyar":"309","Karanganyar Kebumen":"94","Karawang":"95","Kasungan":"96","Kayuagung":"97","Kebumen":"98","Kediri":"99","Kefamenanu":"100","Kendal":"101","Kendari":"102","Kertosono":"103","Ketapang":"104","Kisaran":"105","Klaten":"106","Kolaka":"107","Kota Baru Pulau Laut":"108","Kota Bumi":"109","Kota Jantho":"110","Kota Mobagu":"111","Kuala Kapuas":"112","Kuala Kurun":"113","Kuala Pembuang":"114","Kuala Tungkal":"115","Kudus":"116","Kuningan":"117","Kupang":"118","Kutacane":"119","Kutoarjo":"120","Labuhan":"121","Lahat":"122","Lamongan":"123","Langsa":"124","Larantuka":"125","Lawang":"126","Lhoseumawe":"127","Limboto":"128","Lubuk Basung":"129","Lubuk Linggau":"130","Lubuk Pakam":"131","Lubuk Sikaping":"132","Lumajang":"133","Luwuk":"134","Madiun":"135","Magelang":"136","Magetan":"137","Majalengka":"138","Majene":"139","Makale":"140","Makassar":"141","Malang":"142","Mamuju":"143","Manna":"144","Manokwari":"145","Marabahan":"146","Maros":"147","Martapura":"148","Masohi":"149","Mataram":"150","Maumere":"151","Medan":"152","Mempawah":"153","Menado":"154","Mentok":"155","Merauke":"156","Metro":"157","Meulaboh":"158","Mojokerto":"159","Muara Bulian":"160","Muara Bungo":"161","Muara Enim":"162","Muara Teweh":"163","Muaro Sijunjung":"164","Muntilan":"165","Nabire":"166","Negara":"167","Nganjuk":"168","Ngawi":"169","Nunukan":"170","Pacitan":"171","Padang":"172","Padang Panjang":"173","Padang Sidempuan":"174","Pagaralam":"175","Painan":"176","Palangkaraya":"177","Palembang":"178","Palopo":"179","Palu":"180","Pamekasan":"181","Pandeglang":"182","Pangkajene":"183","Pangkajene Sidenreng":"184","Pangkalanbun":"185","Pangkalpinang":"186","Panyabungan":"187","Pare":"188","Parepare":"189","Pariaman":"190","Pasuruan":"191","Pati":"192","Payakumbuh":"193","Pekalongan":"194","Pekan Baru":"195","Pemalang":"196","Pematangsiantar":"197","Pendopo":"198","Pinrang":"199","Pleihari":"200","Polewali":"201","Pondok Gede":"202","Ponorogo":"203","Pontianak":"204","Poso":"205","Prabumulih":"206","Praya":"207","Probolinggo":"208","Purbalingga":"209","Purukcahu":"210","Purwakarta":"211","Purwodadigrobogan":"212","Purwokerto":"213","Purworejo":"214","Putussibau":"215","Raha":"216","Rangkasbitung":"217","Rantau":"218","Rantauprapat":"219","Rantepao":"220","Rembang":"221","Rengat":"222","Ruteng":"223","Sabang":"224","Salatiga":"225","Samarinda":"226","Sampang":"227","Sampit":"228","Sanggau":"229","Sawahlunto":"230","Sekayu":"231","Selong":"232","Semarang":"233","Sengkang":"234","Serang":"235","Serui":"236","Sibolga":"237","Sidikalang":"238","Sidoarjo":"239","Sigli":"240","Singaparna":"241","Singaraja":"242","Singkawang":"243","Sinjai":"244","Sintang":"245","Situbondo":"246","Slawi":"247","Sleman":"248","Soasiu":"249","Soe":"250","Solo":"251","Solok":"252","Soreang":"253","Sorong":"254","Sragen":"255","Stabat":"256","Subang":"257","Sukabumi":"258","Sukoharjo":"259","Sumbawa Besar":"260","Sumedang":"261","Sumenep":"262","Sungai Liat":"263","Sungai Penuh":"264","Sungguminasa":"265","Surabaya":"266","Surakarta":"267","Tabanan":"268","Tahuna":"269","Takalar":"270","Takengon":"271","Tamiang Layang":"272","Tanah Grogot":"273","Tangerang":"274","Tanjung Balai":"275","Tanjung Enim":"276","Tanjung Pandan":"277","Tanjung Pinang":"278","Tanjung Redep":"279","Tanjung Selor":"280","Tapak Tuan":"281","Tarakan":"282","Tarutung":"283","Tasikmalaya":"284","Tebing Tinggi":"285","Tegal":"286","Temanggung":"287","Tembilahan":"288","Tenggarong":"289","Ternate":"290","Tolitoli":"291","Tondano":"292","Trenggalek":"293","Tual":"294","Tuban":"295","Tulung Agung":"296","Ujung Berung":"297","Ungaran":"298","Waikabubak":"299","Waingapu":"300","Wamena":"301","Watampone":"302","Watansoppeng":"303","Wates":"304","Wonogiri":"305","Wonosari":"306","Wonosobo":"307","Yogyakarta":"308"}`)

func main() {
	app := cli.NewApp()
	app.Name = "go-salat"
	app.Usage = "Salat time for Indonesia"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lokasi, l",
			Value: "Jakarta",
			Usage: "Ambil lokasi",
		},
	}
	app.Version = "1.0.0"

	app.Action = func(c *cli.Context) error {
		inputValue := strings.Title(strings.Replace(c.String("lokasi"), "-", " ", -1))
		if c.NArg() > 0 {
			inputValue = c.Args().Get(0)
		}
		city := make(map[string]string)
		var id string

		err := json.Unmarshal(j, &city)
		if err != nil {
			log.Fatal(err)
		}

		for n, v := range city {
			if n == inputValue {
				id = v
			}
		}

		if id == "" {
			id = "83"
		}

		url := "http://jadwalsholat.pkpu.or.id/monthly.php?id=" + id

		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}

		month := doc.Find("table tr.table_title td b").First().Text()
		//Make sure inputValue always return the correct city name
		inputValue = doc.Find("table tr.table_title td b").Last().Text()

		doc.Find("table tr.table_highlight td").Each(func(i int, s *goquery.Selection) {
			date := s.First().Text()
			var time string

			switch i {
			case 0:
				color.Blue("Waktu salat untuk wilayah: %s", inputValue)
				color.Blue("Tanggal: %s %s", date, month)
				color.Blue("-------------")
			case 1:
				time = s.Empty().Text()
				color.Red("Shubuh: %s", time)
			case 2:
				time = s.Empty().Text()
				color.White("Dzuhur: %s", time)
			case 3:
				time = s.Empty().Text()
				color.Green("Ashr: %s", time)
			case 4:
				time = s.Empty().Text()
				color.Cyan("Maghrib: %s", time)
			case 5:
				time = s.Empty().Text()
				color.White("isya: %s", time)
			}
		})

		return nil
	}

	app.Run(os.Args)

}
