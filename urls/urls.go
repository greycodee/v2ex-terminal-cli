package urls

const baseWebUrl = "https://v2ex.com/"

type TabUrl string


var TabLinkList = []map[string]string{
	{
		"tabName":"技术",
		"link":TabTech,
	},
	{
		"tabName":"创意",
		"link":TabCreative,
	},
	{
		"tabName":"好玩",
		"link":TabPlay,
	},
	{
		"tabName":"Apple",
		"link":TabApple,
	},
	{
		"tabName":"酷工作",
		"link":TabJobs,
	},
	{
		"tabName":"交易",
		"link":TabDeals,
	},
	{
		"tabName":"城市",
		"link":TabCity,
	},
	{
		"tabName":"问与答",
		"link":TabQNA,
	},
	{
		"tabName":"最热",
		"link":TabHot,
	},
	{
		"tabName":"全部",
		"link":TabAll,
	},
}





// tab
const (
	// TabTech 技术
	TabTech = baseWebUrl +"?tab=tech"

	// TabCreative 创意
	TabCreative = baseWebUrl +"?tab=creative"

	// TabPlay 好玩
	TabPlay = baseWebUrl +"?tab=play"

	// TabApple Apple
	TabApple = baseWebUrl +"?tab=apple"

	// TabJobs 酷工作
	TabJobs = baseWebUrl +"?tab=jobs"

	// TabDeals 交易
	TabDeals = baseWebUrl +"?tab=deals"

	// TabCity 城市
	TabCity = baseWebUrl +"?tab=city"

	// TabQNA 问与答
	TabQNA = baseWebUrl +"?tab=qna"

	// TabHot 最热
	TabHot = baseWebUrl +"?tab=hot"

	// TabAll 全部
	TabAll = baseWebUrl +"?tab=all"
)

type url struct {

}
