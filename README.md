# go-lambda-test

## 手動デプロイ方法
### 1. コンパイル
```
GOOS=linux go build -o 生成するファイル名 コンパイルするファイル名.go
```
### 2. zip化
```
zip 生成するzipファイル名.zip コンパイルで生成した実行ファイル名
```

### 3. AWS Lambdaのダッシュボードへ移動
#### 右上のボタンで関数を作成
![スクリーンショット 2022-07-10 23 41 27](https://user-images.githubusercontent.com/75605907/178149608-67b2269b-8bff-4137-b9a4-2226dbcd2345.png)

### 4. 関数を作成
#### 関数名を設定する
#### ランタイムを `` Go 1.x`` にする。
#### 右下のボタンから作成
![スクリーンショット 2022-07-10 23 43 28](https://user-images.githubusercontent.com/75605907/178149673-1ce2f979-fd35-4ecb-8120-af58a464602c.png)

### 5. 2で生成した.zipファイルをアップロード
![スクリーンショット 2022-07-10 23 46 53](https://user-images.githubusercontent.com/75605907/178149795-246963e2-3972-4e2b-ab6a-ef5a3a4cbaf7.png)

### 6. 設定の関数URLからURLを作成
![スクリーンショット 2022-07-10 23 47 47](https://user-images.githubusercontent.com/75605907/178149836-bb1aa662-60b4-4cba-8a5d-df255c47e62e.png)

