package pkg

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 模拟快递数据
var mockExpressData = map[string]map[string]interface{}{
	"3950055201640": {
		"status":  "200",
		"message": "ok",
		"state":   "3", // 签收
		"data": []map[string]interface{}{
			{
				"time":    "2023-12-01 14:30:00",
				"context": "已签收，签收人：本人",
			},
			{
				"time":    "2023-12-01 10:15:00",
				"context": "【上海市】快递员正在派件（张三，电话：138****1234）",
			},
			{
				"time":    "2023-12-01 08:45:00",
				"context": "【上海市】快件到达上海转运中心",
			},
			{
				"time":    "2023-11-30 20:30:00",
				"context": "【杭州市】快件已发往上海",
			},
			{
				"time":    "2023-11-30 16:20:00",
				"context": "【杭州市】杭州转运中心已发出",
			},
			{
				"time":    "2023-11-30 14:10:00",
				"context": "【杭州市】韵达快递已揽收",
			},
		},
	},
	"773123456789": {
		"status":  "200",
		"message": "ok",
		"state":   "5", // 派件中
		"data": []map[string]interface{}{
			{
				"time":    time.Now().Add(-2 * time.Hour).Format("2006-01-02 15:04:05"),
				"context": "【北京市】快递员正在派件（李四，电话：139****5678）",
			},
			{
				"time":    time.Now().Add(-4 * time.Hour).Format("2006-01-02 15:04:05"),
				"context": "【北京市】快件到达北京转运中心",
			},
			{
				"time":    time.Now().Add(-8 * time.Hour).Format("2006-01-02 15:04:05"),
				"context": "【广州市】快件已发往北京",
			},
			{
				"time":    time.Now().Add(-12 * time.Hour).Format("2006-01-02 15:04:05"),
				"context": "【广州市】广州转运中心已发出",
			},
			{
				"time":    time.Now().Add(-16 * time.Hour).Format("2006-01-02 15:04:05"),
				"context": "【广州市】申通快递已揽收",
			},
		},
	},
}

func Search() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== 快递模拟查询demo ===")
	fmt.Println("支持的单号示例：3950055201640（已签收）, 773123456789（派件中）")
	fmt.Print("请输入快递单号: ")
	trackingNum, _ := reader.ReadString('\n')
	trackingNum = strings.TrimSpace(trackingNum)

	fmt.Print("请输入快递公司编码(如yunda、sto、zto等): ")
	companyCode, _ := reader.ReadString('\n')
	companyCode = strings.TrimSpace(companyCode)

	// 模拟网络延迟
	fmt.Printf("\n正在查询 %s 快递单号: %s\n", companyCode, trackingNum)
	fmt.Print("查询中")
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()

	// 检查是否有模拟数据
	if data, exists := mockExpressData[trackingNum]; exists {
		fmt.Println("\n=== 查询结果 ===")
		fmt.Println("查询成功!")

		// 显示物流轨迹
		if trackData, ok := data["data"].([]map[string]interface{}); ok {
			fmt.Println("\n物流轨迹:")
			for i, item := range trackData {
				time := item["time"].(string)
				context := item["context"].(string)
				fmt.Printf("%d. [%s] %s\n", i+1, time, context)
			}
		}

		// 显示快递状态
		if state, ok := data["state"].(string); ok {
			states := map[string]string{
				"0": "在途",
				"1": "揽收",
				"2": "疑难",
				"3": "签收",
				"4": "退签",
				"5": "派件",
				"6": "退回",
			}
			fmt.Printf("\n当前状态: %s\n", states[state])
		}

		// 显示预计送达时间（模拟）
		if data["state"] == "5" {
			// 如果是派件中，显示预计送达时间
			estimatedTime := time.Now().Add(time.Duration(rand.Intn(4)+2) * time.Hour)
			fmt.Printf("预计送达时间: %s\n", estimatedTime.Format("2006-01-02 15:04"))
		}

	} else {
		// 生成随机模拟数据
		fmt.Println("\n=== 查询结果 ===")
		fmt.Println("查询成功!")

		// 随机生成状态
		states := []string{"0", "1", "3", "5"}
		randomState := states[rand.Intn(len(states))]

		// 生成模拟物流轨迹
		city := ""
		switch companyCode {
		case "yunda":
			city = "上海"
		case "sto":
			city = "北京"
		case "zto":
			city = "杭州"
		case "sf":
			city = "深圳"
		default:
			city = "广州"
		}

		fmt.Println("\n物流轨迹:")
		fmt.Printf("1. [%s] %s转运中心已发出\n",
			time.Now().Add(-24*time.Hour).Format("2006-01-02 15:04:05"), city)
		fmt.Printf("2. [%s] 快件到达%s转运中心\n",
			time.Now().Add(-12*time.Hour).Format("2006-01-02 15:04:05"), city)

		if randomState == "3" || randomState == "5" {
			fmt.Printf("3. [%s] 快递员正在派件\n",
				time.Now().Add(-2*time.Hour).Format("2006-01-02 15:04:05"))
		}

		if randomState == "3" {
			fmt.Printf("4. [%s] 已签收\n",
				time.Now().Add(-1*time.Hour).Format("2006-01-02 15:04:05"))
		}

		// 显示状态
		stateMap := map[string]string{
			"0": "在途",
			"1": "揽收",
			"3": "签收",
			"5": "派件",
		}
		fmt.Printf("\n当前状态: %s\n", stateMap[randomState])
	}

	fmt.Println("\n=== 温馨提示 ===")
	fmt.Println("这是一个模拟查询demo，数据为本地生成")
	fmt.Println("如需真实查询，请申请快递100 API授权")
}
