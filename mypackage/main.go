package main

import (
	"wuliu/pkg"
)

//// 模拟快递数据
//var mockExpressData = map[string]map[string]interface{}{
//	"3950055201640": {
//		"status":  "200",
//		"message": "ok",
//		"state":   "3", // 签收
//		"data": []map[string]interface{}{
//			{
//				"time":    "2023-12-01 14:30:00",
//				"context": "已签收，签收人：本人",
//			},
//			{
//				"time":    "2023-12-01 10:15:00",
//				"context": "【上海市】快递员正在派件（张三，电话：138****1234）",
//			},
//			{
//				"time":    "2023-12-01 08:45:00",
//				"context": "【上海市】快件到达上海转运中心",
//			},
//			{
//				"time":    "2023-11-30 20:30:00",
//				"context": "【杭州市】快件已发往上海",
//			},
//			{
//				"time":    "2023-11-30 16:20:00",
//				"context": "【杭州市】杭州转运中心已发出",
//			},
//			{
//				"time":    "2023-11-30 14:10:00",
//				"context": "【杭州市】韵达快递已揽收",
//			},
//		},
//	},
//	"773123456789": {
//		"status":  "200",
//		"message": "ok",
//		"state":   "5", // 派件中
//		"data": []map[string]interface{}{
//			{
//				"time":    time.Now().Add(-2 * time.Hour).Format("2006-01-02 15:04:05"),
//				"context": "【北京市】快递员正在派件（李四，电话：139****5678）",
//			},
//			{
//				"time":    time.Now().Add(-4 * time.Hour).Format("2006-01-02 15:04:05"),
//				"context": "【北京市】快件到达北京转运中心",
//			},
//			{
//				"time":    time.Now().Add(-8 * time.Hour).Format("2006-01-02 15:04:05"),
//				"context": "【广州市】快件已发往北京",
//			},
//			{
//				"time":    time.Now().Add(-12 * time.Hour).Format("2006-01-02 15:04:05"),
//				"context": "【广州市】广州转运中心已发出",
//			},
//			{
//				"time":    time.Now().Add(-16 * time.Hour).Format("2006-01-02 15:04:05"),
//				"context": "【广州市】申通快递已揽收",
//			},
//		},
//	},
//}

func main() {
	pkg.Search()
}
