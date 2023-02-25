package main

import (
	"fmt"
	"strconv"

	"github.com/takeuchi-shogo/k8s-go-sample/config"
	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/infrastructure/database"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/gateways/repositories"
	"github.com/takeuchi-shogo/k8s-go-sample/usecase/interactor"
)

func main() {
	config := config.NewConfig(".")
	db := database.NewDB(config)

	// maleNames := []string{"大和", "大翔", "隼人", "琉生", "颯太", "陽斗", "悠斗", "慶太", "翔太", "勇気", "誠", "直樹", "健太", "裕太", "拓海", "蓮", "蒼空", "瑛太", "悠真", "拓斗"}
	femaleNames := []string{"葵", "美空", "美咲", "美月", "美緒", "結菜", "結愛", "桜", "優花", "彩花", "翼", "愛菜", "優希", "優衣", "莉子", "琴音", "瑞希", "沙羅", "愛美", "咲希"}
	// maleAges := []int{23, 28, 32, 26, 30, 25, 29, 27, 24, 31, 22, 33, 25, 27, 29, 26, 28, 30, 24, 26}
	// femaleAges := []int{24, 26, 27, 29, 30, 25, 23, 28, 26, 31, 22, 30, 24, 27, 25, 29, 33, 26, 32, 28}
	// maleInterests := []string{"釣り", "料理", "映画鑑賞", "読書", "カフェ巡り", "旅行", "アウトドア", "ランニング", "音楽鑑賞", "筋トレ", "ボードゲーム", "カラオケ", "ゲーム", "写真撮影", "ダンス", "ビール", "ギター", "スノボー", "バスケットボール", "ラーメン"}
	// femaleInterests := []string{"旅行", "読書", "カフェ巡り", "映画鑑賞", "料理", "音楽鑑賞", "美容", "ファッション", "ショッピング", "ダンス", "アウトドア", "スポーツ観戦", "ヨガ", "散歩", "ペット", "美味しいもの", "手芸", "絵画", "料理教室", "音楽フェス"}
	// maleLocations := []string{"北海道", "青森県", "岩手県", "宮城県", "秋田県", "山形県", "福島県", "茨城県", "栃木県", "群馬県", "埼玉県", "千葉県", "東京都", "神奈川県", "新潟県", "富山県", "石川県", "福井県", "山梨県", "長野県"}
	femaleLocations := []string{"青森県", "岩手県", "宮城県", "秋田県", "山形県", "福島県", "茨城県", "栃木県", "群馬県", "埼玉県", "千葉県", "東京都", "神奈川県", "新潟県", "富山県", "石川県", "福井県", "山梨県", "長野県", "岐阜県"}

	interactor := interactor.AccountInteractor{
		AccountRepository:     &repositories.AccountRepository{},
		DBRepository:          &repositories.DBRepository{DB: db},
		UserRepository:        &repositories.UserRepository{},
		UserProfileRepository: &repositories.UserProfileRepository{},
	}
	// Male
	// for i := 0; i < 20; i++ {
	// 	account := &models.Accounts{
	// 		PhoneNumber: "00000000000",
	// 		Email:       "test_" + strconv.Itoa(i) + "@exaple.com",
	// 		Password:    "okokok",
	// 	}
	// 	user := &models.Users{
	// 		DisplayName: maleNames[i],
	// 		Gender:      "M",
	// 		Location:    maleLocations[i],
	// 	}
	// 	_, err := interactor.Signup(account, user)
	// 	fmt.Println(i, err)
	// }

	// Female
	for i := 20; i < 40; i++ {
		account := &models.Accounts{
			PhoneNumber: "00000000000",
			Email:       "test_" + strconv.Itoa(i) + "@exaple.com",
			Password:    "okokok",
		}
		user := &models.Users{
			DisplayName: femaleNames[i-20],
			Gender:      "M",
			Location:    femaleLocations[i-20],
		}
		_, err := interactor.Signup(account, user)
		fmt.Println(i, err)
	}
}
