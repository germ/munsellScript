package main

import (
	"os"
	"fmt"
	"errors"
	"strings"
	"io/ioutil"
	"math"
	"encoding/json"
	"encoding/hex"
)

type Color map[string]string

func main() {
	// Read in the table
	// Forgive me father for I have sinned
	var dearGodICantProgram Color
	json.Unmarshal(MunData, &dearGodICantProgram)

	// Read in the value from stdin
	// My last confession was nine months ago
	req, err := ioutil.ReadAll(os.Stdin)
	if err != nil || len(req) == 0 {
		panic(err)
	}

	//Clean input
	// These are my sins
	input := strings.ToUpper(strings.TrimSpace(string(req)))
	
	//Check if RGB value, process
	if strings.HasPrefix(input, "#") {

		//Check DB
		val, ok := dearGodICantProgram[input]
		if !ok {
			// Not in DB, check for nearest
			var bestMatch string
			bestScore := 255*255*255

			for i, v := range(dearGodICantProgram) {
				diff := hexDiff(input[1:], i[1:])
				if diff < bestScore {
					bestScore = diff
					bestMatch = v
				}
			}
			
			// Print & exit
			fmt.Println(bestMatch)
			return
		}

		// Print & exit
		fmt.Println(val)
		return
	}

	// Munsell value, pull RGB
	for i, v := range(dearGodICantProgram) {
		if v == input {
			fmt.Println(i)
			return
		}
	}

	panic(errors.New("You fucked up Rick, lay off the raviolis."))
}

// hex string difference function, 0 is equal
func hexDiff(h1, h2 string) int {
	first, err := hex.DecodeString(h1)
	second, err := hex.DecodeString(h2)
	if err != nil {
		fmt.Println(first, second)
		panic(err)
	}

	retVal := 0.0
	for i, v := range(first) {
		retVal += math.Abs(float64(v - second[i]))
	}
	
	return int(retVal)
}
var MunData = []byte(`{
"#2B151F":"10RP 1/2",
"#340E22":"10RP 1/4",
"#3C0225":"10RP 1/6",
"#2C151D":"2.5R 1/2",
"#350E1E":"2.5R 1/4",
"#3E0020":"2.5R 1/6",
"#2D151B":"5R 1/2",
"#360E1A":"5R 1/4",
"#2D1519":"7.5R 1/2",
"#370D17":"7.5R 1/4",
"#2D1616":"10R 1/2",
"#380D12":"10R 1/4",
"#2D1613":"2.5YR 1/2",
"#380E0A":"2.5YR 1/4",
"#400202":"2.5YR 1/6",
"#2C170F":"5YR 1/2",
"#2A180C":"7.5YR 1/2",
"#281909":"10YR 1/2",
"#261A07":"2.5Y 1/2",
"#231B07":"5Y 1/2",
"#211C08":"7.5Y 1/2",
"#1E1D0B":"10Y 1/2",
"#1B1E0F":"2.5GY 1/2",
"#191E11":"5GY 1/2",
"#161F14":"7.5GY 1/2",
"#052201":"7.5GY 1/4",
"#141F16":"10GY 1/2",
"#121F18":"2.5G 1/2",
"#111F19":"5G 1/2",
"#101F1A":"7.5G 1/2",
"#0F1F1B":"10G 1/2",
"#0D1F1D":"2.5BG 1/2",
"#0B1F1F":"5BG 1/2",
"#0A1F21":"7.5BG 1/2",
"#091F23":"10BG 1/2",
"#091F25":"2.5B 1/2",
"#091F26":"5B 1/2",
"#0B1E28":"7.5B 1/2",
"#0E1E29":"10B 1/2",
"#121D2A":"2.5PB 1/2",
"#151C2A":"5PB 1/2",
"#0B1C35":"5PB 1/4",
"#1A1B2B":"7.5PB 1/2",
"#181935":"7.5PB 1/4",
"#161741":"7.5PB 1/6",
"#16134B":"7.5PB 1/8",
"#180D54":"7.5PB 1/10",
"#1B045C":"7.5PB 1/12",
"#1E192B":"10PB 1/2",
"#201734":"10PB 1/4",
"#22133E":"10PB 1/6",
"#250E46":"10PB 1/8",
"#28064E":"10PB 1/10",
"#21182A":"2.5P 1/2",
"#251532":"2.5P 1/4",
"#29103A":"2.5P 1/6",
"#2E0942":"2.5P 1/8",
"#241829":"5P 1/2",
"#291330":"5P 1/4",
"#2E0E37":"5P 1/6",
"#33043E":"5P 1/8",
"#261728":"7.5P 1/2",
"#2B122F":"7.5P 1/4",
"#310C35":"7.5P 1/6",
"#271726":"10P 1/2",
"#2D112D":"10P 1/4",
"#330A33":"10P 1/6",
"#281625":"2.5RP 1/2",
"#2F102B":"2.5RP 1/4",
"#350830":"2.5RP 1/6",
"#2A1623":"5RP 1/2",
"#310F28":"5RP 1/4",
"#38062D":"5RP 1/6",
"#2B1521":"7.5RP 1/2",
"#330F25":"7.5RP 1/4",
"#3A0429":"7.5RP 1/6",
"#412B31":"10RP 2/2",
"#4B2532":"10RP 2/4",
"#551D33":"10RP 2/6",
"#5F1035":"10RP 2/8",
"#412B2F":"2.5R 2/2",
"#4D252E":"2.5R 2/4",
"#571D2E":"2.5R 2/6",
"#610E2E":"2.5R 2/8",
"#422B2D":"5R 2/2",
"#4D252A":"5R 2/4",
"#581D28":"5R 2/6",
"#630D25":"5R 2/8",
"#422B2B":"7.5R 2/2",
"#4E2526":"7.5R 2/4",
"#581D21":"7.5R 2/6",
"#62101C":"7.5R 2/8",
"#412C29":"10R 2/2",
"#4E2521":"10R 2/4",
"#581E1A":"10R 2/6",
"#621212":"10R 2/8",
"#412C27":"2.5YR 2/2",
"#4D261C":"2.5YR 2/4",
"#562011":"2.5YR 2/6",
"#402D25":"5YR 2/2",
"#4A2817":"5YR 2/4",
"#532300":"5YR 2/6",
"#3E2E22":"7.5YR 2/2",
"#482A11":"7.5YR 2/4",
"#3C2F21":"10YR 2/2",
"#452C0C":"10YR 2/4",
"#3A3020":"2.5Y 2/2",
"#412E06":"2.5Y 2/4",
"#38301F":"5Y 2/2",
"#3D3001":"5Y 2/4",
"#353120":"7.5Y 2/2",
"#393100":"7.5Y 2/4",
"#333221":"10Y 2/2",
"#343304":"10Y 2/4",
"#303323":"2.5GY 2/2",
"#2E340C":"2.5GY 2/4",
"#2E3325":"5GY 2/2",
"#293512":"5GY 2/4",
"#2B3428":"7.5GY 2/2",
"#21361A":"7.5GY 2/4",
"#153907":"7.5GY 2/6",
"#29342A":"10GY 2/2",
"#1B3720":"10GY 2/4",
"#033A15":"10GY 2/6",
"#27342C":"2.5G 2/2",
"#153825":"2.5G 2/4",
"#26342D":"5G 2/2",
"#103829":"5G 2/4",
"#25342E":"7.5G 2/2",
"#0C382C":"7.5G 2/4",
"#243430":"10G 2/2",
"#07382F":"10G 2/4",
"#233432":"2.5BG 2/2",
"#003833":"2.5BG 2/4",
"#223434":"5BG 2/2",
"#213435":"7.5BG 2/2",
"#213437":"10BG 2/2",
"#213439":"2.5B 2/2",
"#21343A":"5B 2/2",
"#22333C":"7.5B 2/2",
"#043548":"7.5B 2/4",
"#25333D":"10B 2/2",
"#0F344A":"10B 2/4",
"#27323E":"2.5PB 2/2",
"#19334B":"2.5PB 2/4",
"#2A313F":"5PB 2/2",
"#21314C":"5PB 2/4",
"#103159":"5PB 2/6",
"#2F303F":"7.5PB 2/2",
"#2C2F4C":"7.5PB 2/4",
"#292D58":"7.5PB 2/6",
"#252B64":"7.5PB 2/8",
"#242770":"7.5PB 2/10",
"#242379":"7.5PB 2/12",
"#261E83":"7.5PB 2/14",
"#29158D":"7.5PB 2/16",
"#2C0896":"7.5PB 2/18",
"#322F3F":"10PB 2/2",
"#342C4A":"10PB 2/4",
"#352954":"10PB 2/6",
"#372660":"10PB 2/8",
"#3A216A":"10PB 2/10",
"#3D1B72":"10PB 2/12",
"#41117C":"10PB 2/14",
"#352E3E":"2.5P 2/2",
"#392B48":"2.5P 2/4",
"#3D2750":"2.5P 2/6",
"#42225A":"2.5P 2/8",
"#461B63":"2.5P 2/10",
"#4B116C":"2.5P 2/12",
"#382D3D":"5P 2/2",
"#3D2945":"5P 2/4",
"#42254D":"5P 2/6",
"#481F56":"5P 2/8",
"#4E165E":"5P 2/10",
"#530866":"5P 2/12",
"#3A2D3B":"7.5P 2/2",
"#402843":"7.5P 2/4",
"#46244A":"7.5P 2/6",
"#4D1C52":"7.5P 2/8",
"#54115A":"7.5P 2/10",
"#3C2C3A":"10P 2/2",
"#432840":"10P 2/4",
"#4A2246":"10P 2/6",
"#52194E":"10P 2/8",
"#590C54":"10P 2/10",
"#3E2B37":"2.5RP 2/2",
"#46273C":"2.5RP 2/4",
"#4E2041":"2.5RP 2/6",
"#571647":"2.5RP 2/8",
"#5F034D":"2.5RP 2/10",
"#3F2B35":"5RP 2/2",
"#482639":"5RP 2/4",
"#511F3C":"5RP 2/6",
"#5A1341":"5RP 2/8",
"#402B33":"7.5RP 2/2",
"#4A2536":"7.5RP 2/4",
"#531E38":"7.5RP 2/6",
"#5D113B":"7.5RP 2/8",
"#5A4146":"10RP 3/2",
"#683B44":"10RP 3/4",
"#743343":"10RP 3/6",
"#802742":"10RP 3/8",
"#8B1742":"10RP 3/10",
"#5B4143":"2.5R 3/2",
"#693B40":"2.5R 3/4",
"#75333C":"2.5R 3/6",
"#812739":"2.5R 3/8",
"#8C1637":"2.5R 3/10",
"#5B4241":"5R 3/2",
"#6A3B3B":"5R 3/4",
"#753336":"5R 3/6",
"#812930":"5R 3/8",
"#8D182B":"5R 3/10",
"#5B423F":"7.5R 3/2",
"#6A3B37":"7.5R 3/4",
"#75342F":"7.5R 3/6",
"#812A27":"7.5R 3/8",
"#8B1D1E":"7.5R 3/10",
"#5B423D":"10R 3/2",
"#693C33":"10R 3/4",
"#743529":"10R 3/6",
"#7F2D1D":"10R 3/8",
"#892210":"10R 3/10",
"#5A433B":"2.5YR 3/2",
"#673E2E":"2.5YR 3/4",
"#723820":"2.5YR 3/6",
"#7A320E":"2.5YR 3/8",
"#594339":"5YR 3/2",
"#643F2A":"5YR 3/4",
"#6E3B18":"5YR 3/6",
"#574437":"7.5YR 3/2",
"#614126":"7.5YR 3/4",
"#693E10":"7.5YR 3/6",
"#554535":"10YR 3/2",
"#5D4322":"10YR 3/4",
"#644107":"10YR 3/6",
"#524634":"2.5Y 3/2",
"#59451F":"2.5Y 3/4",
"#4F4734":"5Y 3/2",
"#55471D":"5Y 3/4",
"#4D4834":"7.5Y 3/2",
"#51481C":"7.5Y 3/4",
"#4A4934":"10Y 3/2",
"#4C4A1D":"10Y 3/4",
"#474A36":"2.5GY 3/2",
"#464C1F":"2.5GY 3/4",
"#444A38":"5GY 3/2",
"#3F4D24":"5GY 3/4",
"#3B4F0A":"5GY 3/6",
"#404B3C":"7.5GY 3/2",
"#374E2C":"7.5GY 3/4",
"#2D511A":"7.5GY 3/6",
"#3E4C3E":"10GY 3/2",
"#304F32":"10GY 3/4",
"#1F5226":"10GY 3/6",
"#3B4C41":"2.5G 3/2",
"#295039":"2.5G 3/4",
"#065332":"2.5G 3/6",
"#3A4C43":"5G 3/2",
"#24503E":"5G 3/4",
"#394C45":"7.5G 3/2",
"#1F5141":"7.5G 3/4",
"#374C46":"10G 3/2",
"#1B5145":"10G 3/4",
"#374C48":"2.5BG 3/2",
"#165149":"2.5BG 3/4",
"#354C4B":"5BG 3/2",
"#11504E":"5BG 3/4",
"#354C4D":"7.5BG 3/2",
"#0D5053":"7.5BG 3/4",
"#344C4F":"10BG 3/2",
"#0C5057":"10BG 3/4",
"#354C51":"2.5B 3/2",
"#0E4F5C":"2.5B 3/4",
"#354B53":"5B 3/2",
"#144E5F":"5B 3/4",
"#374B55":"7.5B 3/2",
"#1D4D61":"7.5B 3/4",
"#394A56":"10B 3/2",
"#264C63":"10B 3/4",
"#3D4957":"2.5PB 3/2",
"#2F4A65":"2.5PB 3/4",
"#184B72":"2.5PB 3/6",
"#404857":"5PB 3/2",
"#374865":"5PB 3/4",
"#2A4873":"5PB 3/6",
"#15497F":"5PB 3/8",
"#454758":"7.5PB 3/2",
"#414665":"7.5PB 3/4",
"#3E4472":"7.5PB 3/6",
"#3A437E":"7.5PB 3/8",
"#36408B":"7.5PB 3/10",
"#333E97":"7.5PB 3/12",
"#313AA3":"7.5PB 3/14",
"#3134AF":"7.5PB 3/16",
"#332FB9":"7.5PB 3/18",
"#3529C2":"7.5PB 3/20",
"#381ECC":"7.5PB 3/22",
"#3B0DD7":"7.5PB 3/24",
"#494658":"10PB 3/2",
"#4A4364":"10PB 3/4",
"#4B4170":"10PB 3/6",
"#4C3E7B":"10PB 3/8",
"#4E3A87":"10PB 3/10",
"#503590":"10PB 3/12",
"#542F9B":"10PB 3/14",
"#5728A6":"10PB 3/16",
"#5B1FAE":"10PB 3/18",
"#5F10B7":"10PB 3/20",
"#4C4556":"2.5P 3/2",
"#504162":"2.5P 3/4",
"#543E6C":"2.5P 3/6",
"#583A76":"2.5P 3/8",
"#5D3480":"2.5P 3/10",
"#622E89":"2.5P 3/12",
"#672593":"2.5P 3/14",
"#6C199C":"2.5P 3/16",
"#7103A4":"2.5P 3/18",
"#4F4455":"5P 3/2",
"#55405F":"5P 3/4",
"#5B3C68":"5P 3/6",
"#613671":"5P 3/8",
"#67307B":"5P 3/10",
"#6D2783":"5P 3/12",
"#721C8C":"5P 3/14",
"#780894":"5P 3/16",
"#524353":"7.5P 3/2",
"#5A3E5C":"7.5P 3/4",
"#613964":"7.5P 3/6",
"#68336C":"7.5P 3/8",
"#6F2B74":"7.5P 3/10",
"#76207D":"7.5P 3/12",
"#7D0F84":"7.5P 3/14",
"#544351":"10P 3/2",
"#5E3D58":"10P 3/4",
"#66385F":"10P 3/6",
"#6E3066":"10P 3/8",
"#76276E":"10P 3/10",
"#7E1875":"10P 3/12",
"#56424E":"2.5RP 3/2",
"#613C53":"2.5RP 3/4",
"#6B3658":"2.5RP 3/6",
"#742D5D":"2.5RP 3/8",
"#7E2163":"2.5RP 3/10",
"#860D68":"2.5RP 3/12",
"#58424B":"5RP 3/2",
"#653B4E":"5RP 3/4",
"#6F3451":"5RP 3/6",
"#7A2A54":"5RP 3/8",
"#831D58":"5RP 3/10",
"#594248":"7.5RP 3/2",
"#663B49":"7.5RP 3/4",
"#72334A":"7.5RP 3/6",
"#7D284C":"7.5RP 3/8",
"#871A4D":"7.5RP 3/10",
"#725B5E":"10RP 4/2",
"#81555C":"10RP 4/4",
"#8E4D5A":"10RP 4/6",
"#9A4558":"10RP 4/8",
"#A53B57":"10RP 4/10",
"#B12D56":"10RP 4/12",
"#BB1656":"10RP 4/14",
"#735B5C":"2.5R 4/2",
"#825558":"2.5R 4/4",
"#8F4D54":"2.5R 4/6",
"#9C4550":"2.5R 4/8",
"#A73B4D":"2.5R 4/10",
"#B22D4A":"2.5R 4/12",
"#BD1447":"2.5R 4/14",
"#735B5A":"5R 4/2",
"#835553":"5R 4/4",
"#904E4C":"5R 4/6",
"#9D4546":"5R 4/8",
"#A83C40":"5R 4/10",
"#B22F39":"5R 4/12",
"#BE1A33":"5R 4/14",
"#745B58":"7.5R 4/2",
"#83554F":"7.5R 4/4",
"#914E46":"7.5R 4/6",
"#9D463D":"7.5R 4/8",
"#A83D34":"7.5R 4/10",
"#B2322B":"7.5R 4/12",
"#BC2120":"7.5R 4/14",
"#745B56":"10R 4/2",
"#83554A":"10R 4/4",
"#904F3E":"10R 4/6",
"#9B4832":"10R 4/8",
"#A54024":"10R 4/10",
"#AE3811":"10R 4/12",
"#735C53":"2.5YR 4/2",
"#825645":"2.5YR 4/4",
"#8D5136":"2.5YR 4/6",
"#974C26":"2.5YR 4/8",
"#9F470F":"2.5YR 4/10",
"#725C50":"5YR 4/2",
"#80583F":"5YR 4/4",
"#8A542E":"5YR 4/6",
"#92501B":"5YR 4/8",
"#715D4E":"7.5YR 4/2",
"#7D593A":"7.5YR 4/4",
"#865627":"7.5YR 4/6",
"#8C540D":"7.5YR 4/8",
"#6F5E4B":"10YR 4/2",
"#795B36":"10YR 4/4",
"#805920":"10YR 4/6",
"#6C5F4A":"2.5Y 4/2",
"#745E33":"2.5Y 4/4",
"#7A5C1A":"2.5Y 4/6",
"#696049":"5Y 4/2",
"#6F6030":"5Y 4/4",
"#745F14":"5Y 4/6",
"#666149":"7.5Y 4/2",
"#6A612F":"7.5Y 4/4",
"#6D6111":"7.5Y 4/6",
"#61634A":"10Y 4/2",
"#656330":"10Y 4/4",
"#676411":"10Y 4/6",
"#5F634C":"2.5GY 4/2",
"#5F6533":"2.5GY 4/4",
"#5D6617":"2.5GY 4/6",
"#5D644E":"5GY 4/2",
"#586638":"5GY 4/4",
"#546820":"5GY 4/6",
"#586452":"7.5GY 4/2",
"#4F6840":"7.5GY 4/4",
"#456B2D":"7.5GY 4/6",
"#396E13":"7.5GY 4/8",
"#566555":"10GY 4/2",
"#486947":"10GY 4/4",
"#386C38":"10GY 4/6",
"#217029":"10GY 4/8",
"#526558":"2.5G 4/2",
"#406A50":"2.5G 4/4",
"#246E47":"2.5G 4/6",
"#51665B":"5G 4/2",
"#3B6A55":"5G 4/4",
"#176E4F":"5G 4/6",
"#4F665D":"7.5G 4/2",
"#376A59":"7.5G 4/4",
"#056E55":"7.5G 4/6",
"#4E665E":"10G 4/2",
"#346A5D":"10G 4/4",
"#4D6661":"2.5BG 4/2",
"#316A61":"2.5BG 4/4",
"#4C6663":"5BG 4/2",
"#2E6A66":"5BG 4/4",
"#4C6565":"7.5BG 4/2",
"#2E6A6B":"7.5BG 4/4",
"#4C6568":"10BG 4/2",
"#2D696F":"10BG 4/4",
"#4D656A":"2.5B 4/2",
"#306874":"2.5B 4/4",
"#4E646B":"5B 4/2",
"#346777":"5B 4/4",
"#51646C":"7.5B 4/2",
"#3B6679":"7.5B 4/4",
"#186885":"7.5B 4/6",
"#53636D":"10B 4/2",
"#42657B":"10B 4/4",
"#2A6687":"10B 4/6",
"#57626E":"2.5PB 4/2",
"#4A637C":"2.5PB 4/4",
"#3A6489":"2.5PB 4/6",
"#1C6597":"2.5PB 4/8",
"#5A616F":"5PB 4/2",
"#52617C":"5PB 4/4",
"#47628A":"5PB 4/6",
"#376298":"5PB 4/8",
"#2062A5":"5PB 4/10",
"#5E606F":"7.5PB 4/2",
"#5B5F7C":"7.5PB 4/4",
"#565E89":"7.5PB 4/6",
"#525C97":"7.5PB 4/8",
"#4D5BA4":"7.5PB 4/10",
"#4859B0":"7.5PB 4/12",
"#4357BD":"7.5PB 4/14",
"#4054CA":"7.5PB 4/16",
"#3E4FD6":"7.5PB 4/18",
"#3C4AE4":"7.5PB 4/20",
"#3D45EE":"7.5PB 4/22",
"#3E3EF9":"7.5PB 4/24",
"#4036FF":"7.5PB 4/26",
"#615F6F":"10PB 4/2",
"#625D7B":"10PB 4/4",
"#625A88":"10PB 4/6",
"#635894":"10PB 4/8",
"#64559F":"10PB 4/10",
"#6552AB":"10PB 4/12",
"#674EB6":"10PB 4/14",
"#6A49C1":"10PB 4/16",
"#6D43CB":"10PB 4/18",
"#703BD7":"10PB 4/20",
"#7434DF":"10PB 4/22",
"#7729E9":"10PB 4/24",
"#7C16F4":"10PB 4/26",
"#645E6E":"2.5P 4/2",
"#685B7A":"2.5P 4/4",
"#6B5884":"2.5P 4/6",
"#6F548F":"2.5P 4/8",
"#735099":"2.5P 4/10",
"#784BA3":"2.5P 4/12",
"#7C46AD":"2.5P 4/14",
"#813FB7":"2.5P 4/16",
"#8636C1":"2.5P 4/18",
"#8C2BCB":"2.5P 4/20",
"#901ED3":"2.5P 4/22",
"#675D6D":"5P 4/2",
"#6D5A77":"5P 4/4",
"#725681":"5P 4/6",
"#78528A":"5P 4/8",
"#7D4C93":"5P 4/10",
"#83479C":"5P 4/12",
"#893FA5":"5P 4/14",
"#9036AF":"5P 4/16",
"#962BB7":"5P 4/18",
"#9C1AC0":"5P 4/20",
"#6A5C6B":"7.5P 4/2",
"#725873":"7.5P 4/4",
"#79547C":"7.5P 4/6",
"#804E84":"7.5P 4/8",
"#87488C":"7.5P 4/10",
"#8E4193":"7.5P 4/12",
"#96389C":"7.5P 4/14",
"#9D2CA4":"7.5P 4/16",
"#A41AAC":"7.5P 4/18",
"#6C5C68":"10P 4/2",
"#76576F":"10P 4/4",
"#7F5276":"10P 4/6",
"#884B7D":"10P 4/8",
"#904583":"10P 4/10",
"#983C8A":"10P 4/12",
"#A03191":"10P 4/14",
"#A72397":"10P 4/16",
"#6E5C66":"2.5RP 4/2",
"#79566B":"2.5RP 4/4",
"#845070":"2.5RP 4/6",
"#8F4975":"2.5RP 4/8",
"#974179":"2.5RP 4/10",
"#A1377E":"2.5RP 4/12",
"#A92983":"2.5RP 4/14",
"#B21488":"2.5RP 4/16",
"#705B63":"5RP 4/2",
"#7D5565":"5RP 4/4",
"#894E67":"5RP 4/6",
"#95466A":"5RP 4/8",
"#9E3E6C":"5RP 4/10",
"#A8326F":"5RP 4/12",
"#B22172":"5RP 4/14",
"#715B60":"7.5RP 4/2",
"#7F5560":"7.5RP 4/4",
"#8C4E60":"7.5RP 4/6",
"#984561":"7.5RP 4/8",
"#A23C62":"7.5RP 4/10",
"#AD2F63":"7.5RP 4/12",
"#B71B64":"7.5RP 4/14",
"#8A7578":"10RP 5/2",
"#9A6F75":"10RP 5/4",
"#A96872":"10RP 5/6",
"#B66070":"10RP 5/8",
"#C2586E":"10RP 5/10",
"#CF4C6C":"10RP 5/12",
"#D9406B":"10RP 5/14",
"#E52C6A":"10RP 5/16",
"#8B7576":"2.5R 5/2",
"#9B6F71":"2.5R 5/4",
"#AA686C":"2.5R 5/6",
"#B86068":"2.5R 5/8",
"#C55763":"2.5R 5/10",
"#D14C5F":"2.5R 5/12",
"#DC405C":"2.5R 5/14",
"#E72C59":"2.5R 5/16",
"#8B7574":"5R 5/2",
"#9C6F6D":"5R 5/4",
"#AB6865":"5R 5/6",
"#B9615E":"5R 5/8",
"#C65857":"5R 5/10",
"#D24D50":"5R 5/12",
"#DD414A":"5R 5/14",
"#E82F43":"5R 5/16",
"#F3083D":"5R 5/18",
"#8C7572":"7.5R 5/2",
"#9D6F69":"7.5R 5/4",
"#AC695F":"7.5R 5/6",
"#B96154":"7.5R 5/8",
"#C6594A":"7.5R 5/10",
"#D24F40":"7.5R 5/12",
"#DC4436":"7.5R 5/14",
"#E6352A":"7.5R 5/16",
"#F01F1E":"7.5R 5/18",
"#8C7670":"10R 5/2",
"#9D7063":"10R 5/4",
"#AB6A56":"10R 5/6",
"#B86348":"10R 5/8",
"#C45C39":"10R 5/10",
"#CE5427":"10R 5/12",
"#D74C12":"10R 5/14",
"#8C766C":"2.5YR 5/2",
"#9C715E":"2.5YR 5/4",
"#A96C4E":"2.5YR 5/6",
"#B5663D":"2.5YR 5/8",
"#BF6029":"2.5YR 5/10",
"#C65C0E":"2.5YR 5/12",
"#8C766A":"5YR 5/2",
"#997259":"5YR 5/4",
"#A66E46":"5YR 5/6",
"#B06931":"5YR 5/8",
"#B86619":"5YR 5/10",
"#8A7768":"7.5YR 5/2",
"#977454":"7.5YR 5/4",
"#A2703E":"7.5YR 5/6",
"#AA6D26":"7.5YR 5/8",
"#887865":"10YR 5/2",
"#94754F":"10YR 5/4",
"#9D7337":"10YR 5/6",
"#A3711A":"10YR 5/8",
"#867963":"2.5Y 5/2",
"#8F784B":"2.5Y 5/4",
"#967630":"2.5Y 5/6",
"#9B750D":"2.5Y 5/8",
"#827A62":"5Y 5/2",
"#8A7A48":"5Y 5/4",
"#8F792B":"5Y 5/6",
"#807B62":"7.5Y 5/2",
"#857B47":"7.5Y 5/4",
"#887C28":"7.5Y 5/6",
"#7D7C62":"10Y 5/2",
"#7F7D48":"10Y 5/4",
"#817E29":"10Y 5/6",
"#797D64":"2.5GY 5/2",
"#797F4C":"2.5GY 5/4",
"#78802F":"2.5GY 5/6",
"#767E66":"5GY 5/2",
"#728050":"5GY 5/4",
"#6E8338":"5GY 5/6",
"#698516":"5GY 5/8",
"#727E6A":"7.5GY 5/2",
"#698258":"7.5GY 5/4",
"#5F8545":"7.5GY 5/6",
"#55882F":"7.5GY 5/8",
"#498B09":"7.5GY 5/10",
"#6E7F6D":"10GY 5/2",
"#628360":"10GY 5/4",
"#538751":"10GY 5/6",
"#428A41":"10GY 5/8",
"#298D2F":"10GY 5/10",
"#6B8071":"2.5G 5/2",
"#5A8468":"2.5G 5/4",
"#44885F":"2.5G 5/6",
"#218C56":"2.5G 5/8",
"#698074":"5G 5/2",
"#55846D":"5G 5/4",
"#3A8968":"5G 5/6",
"#688075":"7.5G 5/2",
"#528571":"7.5G 5/4",
"#34896E":"7.5G 5/6",
"#678078":"10G 5/2",
"#4F8576":"10G 5/4",
"#2E8974":"10G 5/6",
"#66807A":"2.5BG 5/2",
"#4D857A":"2.5BG 5/4",
"#29897A":"2.5BG 5/6",
"#65807D":"5BG 5/2",
"#4B847F":"5BG 5/4",
"#228882":"5BG 5/6",
"#65807F":"7.5BG 5/2",
"#4A8484":"7.5BG 5/4",
"#1B8888":"7.5BG 5/6",
"#657F81":"10BG 5/2",
"#4B8388":"10BG 5/4",
"#198790":"10BG 5/6",
"#667F83":"2.5B 5/2",
"#4D838C":"2.5B 5/4",
"#1F8696":"2.5B 5/6",
"#687E84":"5B 5/2",
"#51828F":"5B 5/4",
"#2C859B":"5B 5/6",
"#6A7D86":"7.5B 5/2",
"#568092":"7.5B 5/4",
"#3B839F":"7.5B 5/6",
"#6D7D87":"10B 5/2",
"#5D7F94":"10B 5/4",
"#4881A1":"10B 5/6",
"#2382AE":"10B 5/8",
"#707C87":"2.5PB 5/2",
"#647D95":"2.5PB 5/4",
"#557EA3":"2.5PB 5/6",
"#417FB1":"2.5PB 5/8",
"#1C80BF":"2.5PB 5/10",
"#747B88":"5PB 5/2",
"#6C7B95":"5PB 5/4",
"#617BA4":"5PB 5/6",
"#547CB1":"5PB 5/8",
"#437CBF":"5PB 5/10",
"#257DCD":"5PB 5/12",
"#777A88":"7.5PB 5/2",
"#747995":"7.5PB 5/4",
"#6F78A3":"7.5PB 5/6",
"#6B77B1":"7.5PB 5/8",
"#6676BD":"7.5PB 5/10",
"#6174CC":"7.5PB 5/12",
"#5A72DA":"7.5PB 5/14",
"#5370E8":"7.5PB 5/16",
"#4D6DF6":"7.5PB 5/18",
"#4769FF":"7.5PB 5/20",
"#7B7987":"10PB 5/2",
"#7B7794":"10PB 5/4",
"#7B75A2":"10PB 5/6",
"#7B73AE":"10PB 5/8",
"#7C70BA":"10PB 5/10",
"#7D6DC6":"10PB 5/12",
"#7D6AD3":"10PB 5/14",
"#7E66E0":"10PB 5/16",
"#8061EA":"10PB 5/18",
"#835CF6":"10PB 5/20",
"#8655FF":"10PB 5/22",
"#7D7886":"2.5P 5/2",
"#807592":"2.5P 5/4",
"#84729E":"2.5P 5/6",
"#876FAA":"2.5P 5/8",
"#8B6BB4":"2.5P 5/10",
"#9067BF":"2.5P 5/12",
"#9462C9":"2.5P 5/14",
"#995DD4":"2.5P 5/16",
"#9D57DE":"2.5P 5/18",
"#A24FE8":"2.5P 5/20",
"#A746F3":"2.5P 5/22",
"#AC3BFD":"2.5P 5/24",
"#B22BFF":"2.5P 5/26",
"#807785":"5P 5/2",
"#857490":"5P 5/4",
"#8B709A":"5P 5/6",
"#916CA5":"5P 5/8",
"#9668AE":"5P 5/10",
"#9D62B8":"5P 5/12",
"#A25DC1":"5P 5/14",
"#A856CB":"5P 5/16",
"#AE4ED4":"5P 5/18",
"#B445DD":"5P 5/20",
"#BB39E6":"5P 5/22",
"#C127F0":"5P 5/24",
"#C700F9":"5P 5/26",
"#827782":"7.5P 5/2",
"#8B738B":"7.5P 5/4",
"#936E94":"7.5P 5/6",
"#9B699D":"7.5P 5/8",
"#A263A5":"7.5P 5/10",
"#AA5DAE":"7.5P 5/12",
"#B156B6":"7.5P 5/14",
"#B84EBE":"7.5P 5/16",
"#C043C6":"7.5P 5/18",
"#C736CE":"7.5P 5/20",
"#CF20D7":"7.5P 5/22",
"#847681":"10P 5/2",
"#8F7287":"10P 5/4",
"#996C8F":"10P 5/6",
"#A36696":"10P 5/8",
"#AC609C":"10P 5/10",
"#B558A3":"10P 5/12",
"#BD50AA":"10P 5/14",
"#C545B0":"10P 5/16",
"#CE38B7":"10P 5/18",
"#D624BE":"10P 5/20",
"#86767E":"2.5RP 5/2",
"#927183":"2.5RP 5/4",
"#9E6B88":"2.5RP 5/6",
"#AA638D":"2.5RP 5/8",
"#B45D92":"2.5RP 5/10",
"#BE5497":"2.5RP 5/12",
"#C8499C":"2.5RP 5/14",
"#D13DA0":"2.5RP 5/16",
"#DA2CA5":"2.5RP 5/18",
"#E306AB":"2.5RP 5/20",
"#88767C":"5RP 5/2",
"#96707D":"5RP 5/4",
"#A3697F":"5RP 5/6",
"#B06281":"5RP 5/8",
"#BB5A84":"5RP 5/10",
"#C74F86":"5RP 5/12",
"#D14489":"5RP 5/14",
"#DB348B":"5RP 5/16",
"#E51C8E":"5RP 5/18",
"#89767A":"7.5RP 5/2",
"#986F79":"7.5RP 5/4",
"#A66979":"7.5RP 5/6",
"#B46179":"7.5RP 5/8",
"#BF5879":"7.5RP 5/10",
"#CC4D79":"7.5RP 5/12",
"#D6417A":"7.5RP 5/14",
"#E1307B":"7.5RP 5/16",
"#EB117D":"7.5RP 5/18",
"#A49092":"10RP 6/2",
"#B48A8F":"10RP 6/4",
"#C4838B":"10RP 6/6",
"#D07D89":"10RP 6/8",
"#DE7486":"10RP 6/10",
"#EB6B84":"10RP 6/12",
"#F66182":"10RP 6/14",
"#FF5180":"10RP 6/16",
"#FF407E":"10RP 6/18",
"#A59090":"2.5R 6/2",
"#B58A8B":"2.5R 6/4",
"#C58385":"2.5R 6/6",
"#D27C80":"2.5R 6/8",
"#E0747C":"2.5R 6/10",
"#ED6A77":"2.5R 6/12",
"#F96073":"2.5R 6/14",
"#FF516F":"2.5R 6/16",
"#FF406B":"2.5R 6/18",
"#A5908E":"5R 6/2",
"#B68A87":"5R 6/4",
"#C6837F":"5R 6/6",
"#D47C78":"5R 6/8",
"#E27471":"5R 6/10",
"#EF6B6A":"5R 6/12",
"#FA6163":"5R 6/14",
"#FF545B":"5R 6/16",
"#FF4354":"5R 6/18",
"#A6908C":"7.5R 6/2",
"#B78A82":"7.5R 6/4",
"#C68478":"7.5R 6/6",
"#D47D6E":"7.5R 6/8",
"#E27563":"7.5R 6/10",
"#EE6D58":"7.5R 6/12",
"#FA634D":"7.5R 6/14",
"#FF5840":"7.5R 6/16",
"#FF4B33":"7.5R 6/18",
"#A79089":"10R 6/2",
"#B88A7C":"10R 6/4",
"#C6856F":"10R 6/6",
"#D37E62":"10R 6/8",
"#E07853":"10R 6/10",
"#EC7042":"10R 6/12",
"#F6682F":"10R 6/14",
"#FF6113":"10R 6/16",
"#A79086":"2.5YR 6/2",
"#B68B77":"2.5YR 6/4",
"#C58667":"2.5YR 6/6",
"#D18157":"2.5YR 6/8",
"#DC7B45":"2.5YR 6/10",
"#E6762E":"2.5YR 6/12",
"#ED710A":"2.5YR 6/14",
"#A69083":"5YR 6/2",
"#B58C72":"5YR 6/4",
"#C2885E":"5YR 6/6",
"#CD844B":"5YR 6/8",
"#D68034":"5YR 6/10",
"#DD7C14":"5YR 6/12",
"#A59180":"7.5YR 6/2",
"#B28E6D":"7.5YR 6/4",
"#BE8B58":"7.5YR 6/6",
"#C78741":"7.5YR 6/8",
"#CE8524":"7.5YR 6/10",
"#A3927E":"10YR 6/2",
"#AE9068":"10YR 6/4",
"#B88D50":"10YR 6/6",
"#C08B36":"10YR 6/8",
"#C58910":"10YR 6/10",
"#A0937B":"2.5Y 6/2",
"#A99264":"2.5Y 6/4",
"#B19049":"2.5Y 6/6",
"#B88F2B":"2.5Y 6/8",
"#9D947A":"5Y 6/2",
"#A39461":"5Y 6/4",
"#AA9345":"5Y 6/6",
"#AE9322":"5Y 6/8",
"#9A957A":"7.5Y 6/2",
"#9F9660":"7.5Y 6/4",
"#A39642":"7.5Y 6/6",
"#A6961D":"7.5Y 6/8",
"#98967A":"10Y 6/2",
"#9A9760":"10Y 6/4",
"#9C9843":"10Y 6/6",
"#9D991C":"10Y 6/8",
"#94977B":"2.5GY 6/2",
"#939963":"2.5GY 6/4",
"#929B48":"2.5GY 6/6",
"#919C24":"2.5GY 6/8",
"#91987D":"5GY 6/2",
"#8D9B67":"5GY 6/4",
"#899D4F":"5GY 6/6",
"#849F31":"5GY 6/8",
"#8B9981":"7.5GY 6/2",
"#839D6F":"7.5GY 6/4",
"#7AA05C":"7.5GY 6/6",
"#6FA345":"7.5GY 6/8",
"#65A62B":"7.5GY 6/10",
"#879A85":"10GY 6/2",
"#7A9E77":"10GY 6/4",
"#6DA268":"10GY 6/6",
"#5DA558":"10GY 6/8",
"#49A847":"10GY 6/10",
"#2EAB34":"10GY 6/12",
"#839A89":"2.5G 6/2",
"#729F7F":"2.5G 6/4",
"#5EA376":"2.5G 6/6",
"#44A76D":"2.5G 6/8",
"#14AB63":"2.5G 6/10",
"#819B8D":"5G 6/2",
"#6CA086":"5G 6/4",
"#54A480":"5G 6/6",
"#31A87A":"5G 6/8",
"#809B8F":"7.5G 6/2",
"#69A08A":"7.5G 6/4",
"#4EA486":"7.5G 6/6",
"#22A883":"7.5G 6/8",
"#7F9B91":"10G 6/2",
"#66A08F":"10G 6/4",
"#4AA48D":"10G 6/6",
"#0AA88B":"10G 6/8",
"#7E9B93":"2.5BG 6/2",
"#64A093":"2.5BG 6/4",
"#45A492":"2.5BG 6/6",
"#7E9A96":"5BG 6/2",
"#639F98":"5BG 6/4",
"#3FA49B":"5BG 6/6",
"#7E9A99":"7.5BG 6/2",
"#629F9D":"7.5BG 6/4",
"#3CA3A2":"7.5BG 6/6",
"#7E9A9B":"10BG 6/2",
"#639EA2":"10BG 6/4",
"#3CA2AA":"10BG 6/6",
"#80999D":"2.5B 6/2",
"#669DA6":"2.5B 6/4",
"#40A1B0":"2.5B 6/6",
"#82989E":"5B 6/2",
"#6B9CA9":"5B 6/4",
"#4BA0B5":"5B 6/6",
"#85989F":"7.5B 6/2",
"#719BAB":"7.5B 6/4",
"#579DB9":"7.5B 6/6",
"#2DA0C6":"7.5B 6/8",
"#8897A0":"10B 6/2",
"#7899AD":"10B 6/4",
"#639BBB":"10B 6/6",
"#479DC9":"10B 6/8",
"#8B96A1":"2.5PB 6/2",
"#7F97AE":"2.5PB 6/4",
"#7099BD":"2.5PB 6/6",
"#5E9ACB":"2.5PB 6/8",
"#449BD9":"2.5PB 6/10",
"#8E95A1":"5PB 6/2",
"#8695AE":"5PB 6/4",
"#7B96BE":"5PB 6/6",
"#6F96CC":"5PB 6/8",
"#6097DA":"5PB 6/10",
"#4797EA":"5PB 6/12",
"#1798FB":"5PB 6/14",
"#9294A1":"7.5PB 6/2",
"#8E93AE":"7.5PB 6/4",
"#8992BD":"7.5PB 6/6",
"#8592CB":"7.5PB 6/8",
"#8090D8":"7.5PB 6/10",
"#798FE8":"7.5PB 6/12",
"#718DF9":"7.5PB 6/14",
"#9593A0":"10PB 6/2",
"#9492AD":"10PB 6/4",
"#9490BB":"10PB 6/6",
"#948EC8":"10PB 6/8",
"#948BD5":"10PB 6/10",
"#9488E3":"10PB 6/12",
"#9485F1":"10PB 6/14",
"#9581FF":"10PB 6/16",
"#9792A0":"2.5P 6/2",
"#9A90AC":"2.5P 6/4",
"#9D8DB9":"2.5P 6/6",
"#A08AC4":"2.5P 6/8",
"#A486D0":"2.5P 6/10",
"#A882DC":"2.5P 6/12",
"#AC7EE7":"2.5P 6/14",
"#B278F3":"2.5P 6/16",
"#B673FE":"2.5P 6/18",
"#99929F":"5P 6/2",
"#9F8FA9":"5P 6/4",
"#A58BB5":"5P 6/6",
"#AA87BF":"5P 6/8",
"#B083CA":"5P 6/10",
"#B57ED3":"5P 6/12",
"#BC78DE":"5P 6/14",
"#C272E8":"5P 6/16",
"#C86CF2":"5P 6/18",
"#CF63FE":"5P 6/20",
"#9D919C":"7.5P 6/2",
"#A58DA5":"7.5P 6/4",
"#AD88AE":"7.5P 6/6",
"#B584B6":"7.5P 6/8",
"#BE7EC0":"7.5P 6/10",
"#C578C8":"7.5P 6/12",
"#CD72D0":"7.5P 6/14",
"#D56AD9":"7.5P 6/16",
"#DC62E1":"7.5P 6/18",
"#E655EB":"7.5P 6/20",
"#EF47F5":"7.5P 6/22",
"#F734FE":"7.5P 6/24",
"#9F919A":"10P 6/2",
"#A88CA1":"10P 6/4",
"#B487A8":"10P 6/6",
"#BD81AF":"10P 6/8",
"#C77BB6":"10P 6/10",
"#D074BC":"10P 6/12",
"#DA6CC4":"10P 6/14",
"#E263CB":"10P 6/16",
"#EB5AD1":"10P 6/18",
"#F54BD9":"10P 6/20",
"#FE39E1":"10P 6/22",
"#FF1CE8":"10P 6/24",
"#A19098":"2.5RP 6/2",
"#AC8B9C":"2.5RP 6/4",
"#B985A1":"2.5RP 6/6",
"#C47FA6":"2.5RP 6/8",
"#CE79AA":"2.5RP 6/10",
"#D971AF":"2.5RP 6/12",
"#E467B4":"2.5RP 6/14",
"#EF5DB9":"2.5RP 6/16",
"#F850BE":"2.5RP 6/18",
"#FF3EC4":"2.5RP 6/20",
"#FF28C9":"2.5RP 6/22",
"#A29095":"5RP 6/2",
"#B08B97":"5RP 6/4",
"#BE8499":"5RP 6/6",
"#CA7E9A":"5RP 6/8",
"#D6779C":"5RP 6/10",
"#E26E9E":"5RP 6/12",
"#EE63A1":"5RP 6/14",
"#F957A3":"5RP 6/16",
"#FF48A6":"5RP 6/18",
"#FF2EA9":"5RP 6/20",
"#FF03AB":"5RP 6/22",
"#A39094":"7.5RP 6/2",
"#B28A93":"7.5RP 6/4",
"#C18392":"7.5RP 6/6",
"#CD7D91":"7.5RP 6/8",
"#DA7591":"7.5RP 6/10",
"#E76C92":"7.5RP 6/12",
"#F36192":"7.5RP 6/14",
"#FF5393":"7.5RP 6/16",
"#FF4494":"7.5RP 6/18",
"#FF2895":"7.5RP 6/20",
"#BEAAAC":"10RP 7/2",
"#CEA5A9":"10RP 7/4",
"#DE9EA5":"10RP 7/6",
"#ED97A2":"10RP 7/8",
"#FB8F9F":"10RP 7/10",
"#FF859C":"10RP 7/12",
"#FF7A99":"10RP 7/14",
"#FF6E97":"10RP 7/16",
"#BFAAAA":"2.5R 7/2",
"#CFA4A5":"2.5R 7/4",
"#E09E9F":"2.5R 7/6",
"#EF979A":"2.5R 7/8",
"#FD8F94":"2.5R 7/10",
"#FF858E":"2.5R 7/12",
"#FF7A8A":"2.5R 7/14",
"#FF6E85":"2.5R 7/16",
"#BFAAA8":"5R 7/2",
"#D1A4A1":"5R 7/4",
"#E19E99":"5R 7/6",
"#F19791":"5R 7/8",
"#FF8F8A":"5R 7/10",
"#FF8581":"5R 7/12",
"#FF7B7A":"5R 7/14",
"#C0AAA6":"7.5R 7/2",
"#D2A49C":"7.5R 7/4",
"#E29E92":"7.5R 7/6",
"#F29786":"7.5R 7/8",
"#FF907C":"7.5R 7/10",
"#FF8770":"7.5R 7/12",
"#FF7E64":"7.5R 7/14",
"#FF7456":"7.5R 7/16",
"#C1AAA4":"10R 7/2",
"#D2A596":"10R 7/4",
"#E29F89":"10R 7/6",
"#F1997A":"10R 7/8",
"#FE936C":"10R 7/10",
"#FF8B5B":"10R 7/12",
"#FF8448":"10R 7/14",
"#FF7B31":"10R 7/16",
"#C1AAA1":"2.5YR 7/2",
"#D2A691":"2.5YR 7/4",
"#E1A080":"2.5YR 7/6",
"#EE9B6F":"2.5YR 7/8",
"#FA965E":"2.5YR 7/10",
"#FF9048":"2.5YR 7/12",
"#FF8B2B":"2.5YR 7/14",
"#C1AB9D":"5YR 7/2",
"#D0A78B":"5YR 7/4",
"#DDA377":"5YR 7/6",
"#E99F64":"5YR 7/8",
"#F39B4F":"5YR 7/10",
"#FC9633":"5YR 7/12",
"#C0AC9A":"7.5YR 7/2",
"#CDA886":"7.5YR 7/4",
"#D9A570":"7.5YR 7/6",
"#E3A25A":"7.5YR 7/8",
"#EB9F41":"7.5YR 7/10",
"#F39C1C":"7.5YR 7/12",
"#BDAD97":"10YR 7/2",
"#C9AA80":"10YR 7/4",
"#D3A868":"10YR 7/6",
"#DBA650":"10YR 7/8",
"#E2A433":"10YR 7/10",
"#BAAE94":"2.5Y 7/2",
"#C4AC7C":"2.5Y 7/4",
"#CDAB63":"2.5Y 7/6",
"#D3AA48":"2.5Y 7/8",
"#D9A822":"2.5Y 7/10",
"#B7AF93":"5Y 7/2",
"#BEAF79":"5Y 7/4",
"#C4AE5E":"5Y 7/6",
"#C9AE40":"5Y 7/8",
"#CDAD0E":"5Y 7/10",
"#B5B093":"7.5Y 7/2",
"#B9B078":"7.5Y 7/4",
"#BDB15C":"7.5Y 7/6",
"#C0B13D":"7.5Y 7/8",
"#B2B193":"10Y 7/2",
"#B5B278":"10Y 7/4",
"#B7B35D":"10Y 7/6",
"#B8B43C":"10Y 7/8",
"#AFB294":"2.5GY 7/2",
"#AEB47B":"2.5GY 7/4",
"#ADB561":"2.5GY 7/6",
"#ACB742":"2.5GY 7/8",
"#ABB80B":"2.5GY 7/10",
"#ABB295":"5GY 7/2",
"#A7B57F":"5GY 7/4",
"#A4B868":"5GY 7/6",
"#A0BA4D":"5GY 7/8",
"#9BBC26":"5GY 7/10",
"#A5B49A":"7.5GY 7/2",
"#9DB787":"7.5GY 7/4",
"#94BB74":"7.5GY 7/6",
"#8BBE5F":"7.5GY 7/8",
"#81C146":"7.5GY 7/10",
"#76C324":"7.5GY 7/12",
"#A1B49E":"10GY 7/2",
"#94B98F":"10GY 7/4",
"#88BD80":"10GY 7/6",
"#79C070":"10GY 7/8",
"#67C45F":"10GY 7/10",
"#51C74D":"10GY 7/12",
"#32CA37":"10GY 7/14",
"#9DB5A3":"2.5G 7/2",
"#8CBA98":"2.5G 7/4",
"#79BE8F":"2.5G 7/6",
"#62C284":"2.5G 7/8",
"#43C67A":"2.5G 7/10",
"#9BB5A6":"5G 7/2",
"#86BA9F":"5G 7/4",
"#70BF99":"5G 7/6",
"#53C393":"5G 7/8",
"#18C78D":"5G 7/10",
"#99B5A9":"7.5G 7/2",
"#84BBA4":"7.5G 7/4",
"#6BBF9F":"7.5G 7/6",
"#4AC39C":"7.5G 7/8",
"#98B5AB":"10G 7/2",
"#81BBA8":"10G 7/4",
"#67BFA6":"10G 7/6",
"#40C3A4":"10G 7/8",
"#98B5AD":"2.5BG 7/2",
"#80BAAC":"2.5BG 7/4",
"#64BFAC":"2.5BG 7/6",
"#36C3AB":"2.5BG 7/8",
"#97B5B0":"5BG 7/2",
"#7EBAB2":"5BG 7/4",
"#61BFB4":"5BG 7/6",
"#29C3B6":"5BG 7/8",
"#97B5B3":"7.5BG 7/2",
"#7DBAB7":"7.5BG 7/4",
"#5FBEBB":"7.5BG 7/6",
"#21C2BF":"7.5BG 7/8",
"#98B4B5":"10BG 7/2",
"#7DB9BC":"10BG 7/4",
"#5FBDC2":"10BG 7/6",
"#1DC1C9":"10BG 7/8",
"#9AB4B7":"2.5B 7/2",
"#80B8C0":"2.5B 7/4",
"#61BCC9":"2.5B 7/6",
"#28C0D2":"2.5B 7/8",
"#9DB3B8":"5B 7/2",
"#85B7C4":"5B 7/4",
"#68BACF":"5B 7/6",
"#3BBEDB":"5B 7/8",
"#9FB2B9":"7.5B 7/2",
"#8AB6C6":"7.5B 7/4",
"#72B9D2":"7.5B 7/6",
"#4EBBE0":"7.5B 7/8",
"#A2B1BA":"10B 7/2",
"#91B4C7":"10B 7/4",
"#7DB6D5":"10B 7/6",
"#63B8E4":"10B 7/8",
"#35BAF4":"10B 7/10",
"#A6B0BA":"2.5PB 7/2",
"#99B2C9":"2.5PB 7/4",
"#8AB3D7":"2.5PB 7/6",
"#78B5E6":"2.5PB 7/8",
"#5DB6F7":"2.5PB 7/10",
"#A8AFBA":"5PB 7/2",
"#9FB0C9":"5PB 7/4",
"#95B1D8":"5PB 7/6",
"#88B1E8":"5PB 7/8",
"#77B2F9":"5PB 7/10",
"#ACAEBA":"7.5PB 7/2",
"#A7AEC9":"7.5PB 7/4",
"#A3ADD8":"7.5PB 7/6",
"#9DACE8":"7.5PB 7/8",
"#97ABF9":"7.5PB 7/10",
"#AEAEBA":"10PB 7/2",
"#AEACC8":"10PB 7/4",
"#ADAAD6":"10PB 7/6",
"#ADA8E5":"10PB 7/8",
"#ACA6F4":"10PB 7/10",
"#ADA2FF":"10PB 7/12",
"#B1ADB9":"2.5P 7/2",
"#B4AAC6":"2.5P 7/4",
"#B7A8D3":"2.5P 7/6",
"#BAA5E0":"2.5P 7/8",
"#BEA1ED":"2.5P 7/10",
"#C29DFB":"2.5P 7/12",
"#B3ACB8":"5P 7/2",
"#B8A9C4":"5P 7/4",
"#BEA6CF":"5P 7/6",
"#C3A2DA":"5P 7/8",
"#CA9EE6":"5P 7/10",
"#D099F2":"5P 7/12",
"#D693FD":"5P 7/14",
"#B6ACB5":"7.5P 7/2",
"#BFA8BE":"7.5P 7/4",
"#C8A3C7":"7.5P 7/6",
"#D19ED1":"7.5P 7/8",
"#DA98DA":"7.5P 7/10",
"#E293E3":"7.5P 7/12",
"#EB8CEC":"7.5P 7/14",
"#F484F6":"7.5P 7/16",
"#FC7CFF":"7.5P 7/18",
"#B8ABB4":"10P 7/2",
"#C3A7BB":"10P 7/4",
"#CEA1C2":"10P 7/6",
"#D99CC9":"10P 7/8",
"#E495D0":"10P 7/10",
"#ED8FD7":"10P 7/12",
"#F787DE":"10P 7/14",
"#FF7EE6":"10P 7/16",
"#FF74EE":"10P 7/18",
"#FF68F6":"10P 7/20",
"#FF57FE":"10P 7/22",
"#BAABB2":"2.5RP 7/2",
"#C7A6B6":"2.5RP 7/4",
"#D4A0BB":"2.5RP 7/6",
"#E199C0":"2.5RP 7/8",
"#EC93C4":"2.5RP 7/10",
"#F78BC9":"2.5RP 7/12",
"#FF83CE":"2.5RP 7/14",
"#FF78D4":"2.5RP 7/16",
"#FF6BDA":"2.5RP 7/18",
"#FF5BE0":"2.5RP 7/20",
"#BBABB0":"5RP 7/2",
"#CAA5B1":"5RP 7/4",
"#D99FB2":"5RP 7/6",
"#E798B4":"5RP 7/8",
"#F391B5":"5RP 7/10",
"#FF88B7":"5RP 7/12",
"#FF7EBA":"5RP 7/14",
"#FF72BC":"5RP 7/16",
"#FF64BF":"5RP 7/18",
"#BDABAE":"7.5RP 7/2",
"#CCA5AD":"7.5RP 7/4",
"#DC9EAC":"7.5RP 7/6",
"#EA97AB":"7.5RP 7/8",
"#F790AB":"7.5RP 7/10",
"#FF86AB":"7.5RP 7/12",
"#FF7BAB":"7.5RP 7/14",
"#FF6FAC":"7.5RP 7/16",
"#D6C6C7":"10RP 8/2",
"#E9BFC3":"10RP 8/4",
"#FBB9BF":"10RP 8/6",
"#FFB1BB":"10RP 8/8",
"#FFA8B8":"10RP 8/10",
"#D7C6C5":"2.5R 8/2",
"#EBBFBF":"2.5R 8/4",
"#FDB8B9":"2.5R 8/6",
"#FFB1B2":"2.5R 8/8",
"#FFA8AC":"2.5R 8/10",
"#D8C6C4":"5R 8/2",
"#EDBFBA":"5R 8/4",
"#FEB8B2":"5R 8/6",
"#FFB1A9":"5R 8/8",
"#FFA8A1":"5R 8/10",
"#D8C6C2":"7.5R 8/2",
"#EEBFB5":"7.5R 8/4",
"#FFB9AA":"7.5R 8/6",
"#FFB19E":"7.5R 8/8",
"#FFA992":"7.5R 8/10",
"#D9C6BF":"10R 8/2",
"#EFBFB0":"10R 8/4",
"#FFB9A1":"10R 8/6",
"#FFB391":"10R 8/8",
"#FFAD83":"10R 8/10",
"#DAC6BC":"2.5YR 8/2",
"#EEC0A9":"2.5YR 8/4",
"#FEBB99":"2.5YR 8/6",
"#FFB586":"2.5YR 8/8",
"#FFB074":"2.5YR 8/10",
"#FFAA5E":"2.5YR 8/12",
"#DBC6B8":"5YR 8/2",
"#ECC2A4":"5YR 8/4",
"#FABE90":"5YR 8/6",
"#FFB97A":"5YR 8/8",
"#FFB566":"5YR 8/10",
"#FFB14C":"5YR 8/12",
"#FFAD2B":"5YR 8/14",
"#DAC7B4":"7.5YR 8/2",
"#E8C39F":"7.5YR 8/4",
"#F5C089":"7.5YR 8/6",
"#FFBD71":"7.5YR 8/8",
"#FFBA59":"7.5YR 8/10",
"#FFB73B":"7.5YR 8/12",
"#D8C8B1":"10YR 8/2",
"#E4C599":"10YR 8/4",
"#EFC381":"10YR 8/6",
"#F8C168":"10YR 8/8",
"#FFBF4D":"10YR 8/10",
"#FFBD28":"10YR 8/12",
"#D5C9AE":"2.5Y 8/2",
"#DFC895":"2.5Y 8/4",
"#E7C67B":"2.5Y 8/6",
"#EFC561":"2.5Y 8/8",
"#F5C342":"2.5Y 8/10",
"#FAC209":"2.5Y 8/12",
"#D2CAAC":"5Y 8/2",
"#D8CA92":"5Y 8/4",
"#DFCA76":"5Y 8/6",
"#E4C959":"5Y 8/8",
"#E8C935":"5Y 8/10",
"#D0CBAB":"7.5Y 8/2",
"#D4CB91":"7.5Y 8/4",
"#D8CC75":"7.5Y 8/6",
"#DCCC56":"7.5Y 8/8",
"#DECC30":"7.5Y 8/10",
"#CDCCAB":"10Y 8/2",
"#D0CD91":"10Y 8/4",
"#D2CE74":"10Y 8/6",
"#D4CF56":"10Y 8/8",
"#D5CF2F":"10Y 8/10",
"#CACCAC":"2.5GY 8/2",
"#C9CF93":"2.5GY 8/4",
"#C9D078":"2.5GY 8/6",
"#C8D25B":"2.5GY 8/8",
"#C7D335":"2.5GY 8/10",
"#C7CDAE":"5GY 8/2",
"#C4D096":"5GY 8/4",
"#C0D37D":"5GY 8/6",
"#BBD564":"5GY 8/8",
"#B7D743":"5GY 8/10",
"#C0CFB2":"7.5GY 8/2",
"#B7D39F":"7.5GY 8/4",
"#AED68A":"7.5GY 8/6",
"#A6D975":"7.5GY 8/8",
"#9CDC5D":"7.5GY 8/10",
"#91DF3F":"7.5GY 8/12",
"#BAD0B7":"10GY 8/2",
"#AED4A7":"10GY 8/4",
"#A0D897":"10GY 8/6",
"#92DC88":"10GY 8/8",
"#81E076":"10GY 8/10",
"#6DE363":"10GY 8/12",
"#55E64F":"10GY 8/14",
"#2EE936":"10GY 8/16",
"#B6D0BC":"2.5G 8/2",
"#A5D5B1":"2.5G 8/4",
"#90DAA5":"2.5G 8/6",
"#7CDE9B":"2.5G 8/8",
"#61E290":"2.5G 8/10",
"#35E686":"2.5G 8/12",
"#B4D1C0":"5G 8/2",
"#A0D6B9":"5G 8/4",
"#87DBB1":"5G 8/6",
"#6BDFAB":"5G 8/8",
"#42E3A5":"5G 8/10",
"#B2D1C3":"7.5G 8/2",
"#9DD6BD":"7.5G 8/4",
"#82DBB8":"7.5G 8/6",
"#63DFB4":"7.5G 8/8",
"#2FE4B0":"7.5G 8/10",
"#B1D1C5":"10G 8/2",
"#9AD6C2":"10G 8/4",
"#7EDBBF":"10G 8/6",
"#5CE0BD":"10G 8/8",
"#14E4BB":"10G 8/10",
"#B1D0C7":"2.5BG 8/2",
"#98D6C6":"2.5BG 8/4",
"#7BDBC5":"2.5BG 8/6",
"#55DFC5":"2.5BG 8/8",
"#B1D0CA":"5BG 8/2",
"#97D6CB":"5BG 8/4",
"#78DBCD":"5BG 8/6",
"#4BDFCF":"5BG 8/8",
"#B1D0CD":"7.5BG 8/2",
"#97D5D1":"7.5BG 8/4",
"#75DAD6":"7.5BG 8/6",
"#44DFDA":"7.5BG 8/8",
"#B2CFCF":"10BG 8/2",
"#97D5D6":"10BG 8/4",
"#75D9DD":"10BG 8/6",
"#42DEE3":"10BG 8/8",
"#B4CFD1":"2.5B 8/2",
"#98D4DA":"2.5B 8/4",
"#77D8E3":"2.5B 8/6",
"#45DCED":"2.5B 8/8",
"#B8CED3":"5B 8/2",
"#9ED2DF":"5B 8/4",
"#80D6EB":"5B 8/6",
"#4FDAF9":"5B 8/8",
"#BACDD3":"7.5B 8/2",
"#A3D1E1":"7.5B 8/4",
"#88D4EF":"7.5B 8/6",
"#60D8FF":"7.5B 8/8",
"#BDCCD4":"10B 8/2",
"#ABCFE3":"10B 8/4",
"#95D2F2":"10B 8/6",
"#75D4FF":"10B 8/8",
"#C1CBD4":"2.5PB 8/2",
"#B2CDE4":"2.5PB 8/4",
"#A1CFF5":"2.5PB 8/6",
"#C3CAD4":"5PB 8/2",
"#B9CBE4":"5PB 8/4",
"#ACCCF6":"5PB 8/6",
"#C7C9D4":"7.5PB 8/2",
"#C1C9E4":"7.5PB 8/4",
"#BBC8F6":"7.5PB 8/6",
"#C9C9D3":"10PB 8/2",
"#C8C7E3":"10PB 8/4",
"#C7C5F3":"10PB 8/6",
"#C7C2FF":"10PB 8/8",
"#CBC8D3":"2.5P 8/2",
"#CEC5E1":"2.5P 8/4",
"#D1C3F0":"2.5P 8/6",
"#D4BFFF":"2.5P 8/8",
"#CDC8D2":"5P 8/2",
"#D2C4DF":"5P 8/4",
"#D8C1EB":"5P 8/6",
"#DEBCF9":"5P 8/8",
"#E5B8FF":"5P 8/10",
"#D0C7CF":"7.5P 8/2",
"#DAC2D9":"7.5P 8/4",
"#E3BEE2":"7.5P 8/6",
"#EEB8ED":"7.5P 8/8",
"#F7B2F7":"7.5P 8/10",
"#FFACFF":"7.5P 8/12",
"#D1C7CE":"10P 8/2",
"#DEC2D5":"10P 8/4",
"#EABCDC":"10P 8/6",
"#F6B6E4":"10P 8/8",
"#FFAFEC":"10P 8/10",
"#FFA8F4":"10P 8/12",
"#FF9FFC":"10P 8/14",
"#D3C6CC":"2.5RP 8/2",
"#E1C1D0":"2.5RP 8/4",
"#F0BAD5":"2.5RP 8/6",
"#FEB4DA":"2.5RP 8/8",
"#FFACDE":"2.5RP 8/10",
"#FFA4E4":"2.5RP 8/12",
"#FF9AE9":"2.5RP 8/14",
"#D4C6CA":"5RP 8/2",
"#E5C0CB":"5RP 8/4",
"#F5B9CC":"5RP 8/6",
"#FFB2CE":"5RP 8/8",
"#FFAAD0":"5RP 8/10",
"#FFA0D2":"5RP 8/12",
"#D5C6C8":"7.5RP 8/2",
"#E7C0C7":"7.5RP 8/4",
"#F8B9C6":"7.5RP 8/6",
"#FFB1C5":"7.5RP 8/8",
"#FFA9C4":"7.5RP 8/10",
"#FF9EC4":"7.5RP 8/12",
"#F1E1E2":"10RP 9/2",
"#FFDADD":"10RP 9/4",
"#FFD2D9":"10RP 9/6",
"#F1E1E1":"2.5R 9/2",
"#FFDAD9":"2.5R 9/4",
"#FFD2D2":"2.5R 9/6",
"#F3E1DF":"5R 9/2",
"#FFDAD4":"5R 9/4",
"#FFD2CA":"5R 9/6",
"#F4E1DD":"7.5R 9/2",
"#FFDACE":"7.5R 9/4",
"#FFD3C1":"7.5R 9/6",
"#F5E1DA":"10R 9/2",
"#FFDAC8":"10R 9/4",
"#FFD4B8":"10R 9/6",
"#F6E1D6":"2.5YR 9/2",
"#FFDBC2":"2.5YR 9/4",
"#FFD6AF":"2.5YR 9/6",
"#F6E1D2":"5YR 9/2",
"#FFDDBB":"5YR 9/4",
"#FFD8A6":"5YR 9/6",
"#F6E2CD":"7.5YR 9/2",
"#FFDEB5":"7.5YR 9/4",
"#FFDB9F":"7.5YR 9/6",
"#FFD887":"7.5YR 9/8",
"#F4E3C9":"10YR 9/2",
"#FFE1B0":"10YR 9/4",
"#FFDF97":"10YR 9/6",
"#FFDC7D":"10YR 9/8",
"#F1E4C7":"2.5Y 9/2",
"#FBE3AC":"2.5Y 9/4",
"#FFE292":"2.5Y 9/6",
"#FFE076":"2.5Y 9/8",
"#FFDF59":"2.5Y 9/10",
"#FFDE33":"2.5Y 9/12",
"#EEE5C5":"5Y 9/2",
"#F4E6A9":"5Y 9/4",
"#FAE58D":"5Y 9/6",
"#FFE570":"5Y 9/8",
"#FFE550":"5Y 9/10",
"#FFE41D":"5Y 9/12",
"#EBE6C4":"7.5Y 9/2",
"#F0E7A8":"7.5Y 9/4",
"#F4E88B":"7.5Y 9/6",
"#F7E86D":"7.5Y 9/8",
"#FAE84B":"7.5Y 9/10",
"#FDE80B":"7.5Y 9/12",
"#E9E7C4":"10Y 9/2",
"#ECE8A8":"10Y 9/4",
"#EEEA8B":"10Y 9/6",
"#F0EA6C":"10Y 9/8",
"#F1EB4A":"10Y 9/10",
"#E6E8C4":"2.5GY 9/2",
"#E5EAA9":"2.5GY 9/4",
"#E5EC8D":"2.5GY 9/6",
"#E4EE6F":"2.5GY 9/8",
"#E3EF4E":"2.5GY 9/10",
"#E2F10D":"2.5GY 9/12",
"#E2E9C6":"5GY 9/2",
"#DFECAB":"5GY 9/4",
"#DBEF91":"5GY 9/6",
"#D7F175":"5GY 9/8",
"#D3F356":"5GY 9/10",
"#CFF528":"5GY 9/12",
"#DAEBCB":"7.5GY 9/2",
"#D2EFB5":"7.5GY 9/4",
"#C8F39E":"7.5GY 9/6",
"#BFF688":"7.5GY 9/8",
"#B5F96F":"7.5GY 9/10",
"#ABFC52":"7.5GY 9/12",
"#A1FE2B":"7.5GY 9/14",
"#D5ECD0":"10GY 9/2",
"#C7F0BF":"10GY 9/4",
"#B7F5AC":"10GY 9/6",
"#A9F99C":"10GY 9/8",
"#97FD89":"10GY 9/10",
"#85FF77":"10GY 9/12",
"#6EFF63":"10GY 9/14",
"#4FFF4A":"10GY 9/16",
"#0FFF28":"10GY 9/18",
"#D0ECD6":"2.5G 9/2",
"#BDF2CA":"2.5G 9/4",
"#A7F7BC":"2.5G 9/6",
"#91FCB1":"2.5G 9/8",
"#76FFA5":"2.5G 9/10",
"#54FF9B":"2.5G 9/12",
"#CDECDA":"5G 9/2",
"#B7F2D2":"5G 9/4",
"#9CF8C9":"5G 9/6",
"#7EFDC3":"5G 9/8",
"#58FFBD":"5G 9/10",
"#CCEDDD":"7.5G 9/2",
"#B4F3D7":"7.5G 9/4",
"#96F8D1":"7.5G 9/6",
"#74FDCC":"7.5G 9/8",
"#47FFC8":"7.5G 9/10",
"#CBEDE0":"10G 9/2",
"#B1F3DC":"10G 9/4",
"#92F8D9":"10G 9/6",
"#6BFDD6":"10G 9/8",
"#32FFD4":"10G 9/10",
"#CAECE2":"2.5BG 9/2",
"#AFF3E1":"2.5BG 9/4",
"#8EF8DF":"2.5BG 9/6",
"#63FDDF":"2.5BG 9/8",
"#0FFFDE":"2.5BG 9/10",
"#CAECE5":"5BG 9/2",
"#AEF2E6":"5BG 9/4",
"#8AF8E7":"5BG 9/6",
"#59FDE9":"5BG 9/8",
"#CBECE8":"7.5BG 9/2",
"#ADF2EC":"7.5BG 9/4",
"#88F8F1":"7.5BG 9/6",
"#4FFDF5":"7.5BG 9/8",
"#CCEBEB":"10BG 9/2",
"#ADF1F2":"10BG 9/4",
"#87F7F9":"10BG 9/6",
"#CFEAED":"2.5B 9/2",
"#AFF0F7":"2.5B 9/4",
"#D2E9EE":"5B 9/2",
"#B3EFFC":"5B 9/4",
"#D5E8EF":"7.5B 9/2",
"#B9EDFF":"7.5B 9/4",
"#D8E8EF":"10B 9/2",
"#C0EBFF":"10B 9/4",
"#DCE6EF":"2.5PB 9/2",
"#DEE6EF":"5PB 9/2",
"#E2E5EF":"7.5PB 9/2",
"#E5E4EE":"10PB 9/2",
"#E3E2FF":"10PB 9/4",
"#E6E3EE":"2.5P 9/2",
"#E9E1FE":"2.5P 9/4",
"#E8E3ED":"5P 9/2",
"#EEDFFD":"5P 9/4",
"#EBE2EA":"7.5P 9/2",
"#F7DDF5":"7.5P 9/4",
"#FFD8FF":"7.5P 9/6",
"#EDE2E9":"10P 9/2",
"#FBDCF1":"10P 9/4",
"#FFD6F9":"10P 9/6",
"#EEE2E7":"2.5RP 9/2",
"#FFDCEC":"2.5RP 9/4",
"#FFD5F1":"2.5RP 9/6",
"#F0E1E5":"5RP 9/2",
"#FFDBE6":"5RP 9/4",
"#FFD3E7":"5RP 9/6",
"#F1E1E4":"7.5RP 9/2",
"#FFDAE2":"7.5RP 9/4",
"#FFD3E0":"7.5RP 9/6",
"#222221":"N1",
"#3B3A3A":"N2",
"#525251":"N3",
"#6B6C6B":"N4",
"#888987":"N5",
"#A3A2A2":"N6",
"#B5B6B5":"N7",
"#CACACA":"N8",
"#E9E8E7":"N9"
}`)
