//Package dice 简单骰子--a 公牛大粪-like gift for myrnfc
package dice

import (
	"math/rand"
	"strconv"
	"strings"

	control "github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := control.Register("dice", &control.Options{
		DisableOnDefault: false,
		Help: "dice\n" +
			"- [ ! | ！ ]r[骰子数量]d[骰子面数]\n",
	})
	engine.OnRegex("[!！][rR](.*)[dD](.*)").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			r1 := ctx.State["regex_matched"].([]string)[1]
			d1 := ctx.State["regex_matched"].([]string)[2]
			if r1 == "" {
				r1 = "1"
			}
			if d1 == "" {
				d1 = "100"
			}
			r, _ := strconv.Atoi(r1)
			d, _ := strconv.Atoi(d1)
			if r < 1 || d <= 1 {
				ctx.SendChain(message.Text("阁下..你在让我骰什么啊？( ´_ゝ`)"))
				return
			}
			if r <= 100 && d <= 100 {
				res, sum := rd(r, d)
				ctx.SendChain(message.At(ctx.Event.UserID), message.Text(" 阁下掷出了", "R", r, "D", d, "=", sum, "\n", res, sum, "呢~"))
			} else {
				ctx.SendChain(message.Text("骰子太多啦~~数不过来了！"))
			}
		})
}

func rd(r, d int) (string, int) {
	var res string
	var sum int
	for i := 0; i < r; i++ {
		sum += rand.Intn(d-1) + 1
		res += strconv.Itoa(sum) + "+"
	}
	res += "="
	res = strings.ReplaceAll(res, "+=", "=")
	return res, sum
}
