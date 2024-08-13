# 개요

웹 페이지를 통해 `동영상 자르기`, `동영상 붙이기` 작업을 수행하는 웹 서버

<br>

# 사용예시

1. 프로젝트 실행하여 웹서버 구동
2. 브라우저로 [`localhost:3000`](http://localhost:3000) 접속
3. 좌측 `동영상 업로드` 영역

   -  `파일 선택` 버튼을 통해 업로드할 동영상 파일 선택 후 `업로드` 버튼 클릭하여 파일 업로드
      -  `.avi`, `.mp4`, `.mov` 만 허용
   -  `./tempvideo` 경로에 있는 테스트용 동영상 파일 활용가능

4. 중앙 `동영상 편집` 영역

   -  `동영상 붙이기` 작업을 할 경우 사용되는 동영상 수에 맞게 `추가하기` 버튼을 클릭하여 행 추가
   -  `동영상 붙이기` 작업을 할 경우 원하는 순서대로 콤보박스에서 선택
   -  `동영상 자르기` 작업을 할 경우 `첫 번째 입력칸: 시작시간`, `두 번째 입력칸: 끝시간` 을 입력
      -  `시작시간` 혹은 `끝시간`을 입력하지 않을 경우 `원본 동영상`의 `시작지점`과 `종료지점`으로 설정됨
   -  위 설정값 입력 후 `동영상 생성` 버튼 클릭시 동영상이 생성되고 `다운로드 버튼`, `다운로드 URL링크` 화면 출력

5. 우측 `동영상 조회` 영역

   -  `조회하기` 버튼 클릭시 `업로드된 동영상`, `생성된 동영상` 목록 출력
   -  `업로드된 동영상` 표시 항목: `업로드 동영상ID`, `업로드 일시`, `원본 파일명`, `업로드 파일 경로`
   -  `생성된 동영상` 표시 항목: `생성된 동영상ID`, `생성 일시`, `요청 내역`, `생성된 파일 경로`

<br>

# End point

### `POST`

`/upload` - 파일 업로드  
`/create_video` - 동영상 생성

### `GET`

`/download/:filename` - 생성된 동영상 다운로드  
`/get_video_list` - 업로드 및 생성된 동영상 목록 조회

<br>

# 프로젝트 실행시 필요 항목

-  ## `GO`, `ffmpeg`, `MariaDB` install & `MariaDB` start

   ### MAC

   `homebrew` 설치 https://brew.sh/ko/

   ```shell
   brew install go
   brew install ffmpeg
   brew install mariadb
   brew services start mariadb
   ```

   ### Windows

   `chocolatey` 설치 https://chocolatey.org/install

   ```shell
   choco install golang
   choco install ffmpeg
   choco install mariadb
   net start mariadb

   ```

-  ## set schema

   `mariadb` 접속 후 `videdit` 프로젝트 경로의 `schema.sql` 쿼리 실행

<br>

# 프로젝트 구성

`config` - 설정  
`ffmpeg` - 동영상 처리  
`handler` - http 요청 처리  
`mariadb` - database 로직 처리  
`model` - entity 모음  
`router` - endpoint 라우팅  
`util` - 파일 처리, Error Response 관련 유틸

<br>

# 의존성

-  [go mysql driver](https://github.com/go-sql-driver/mysql) , [gorm](https://gorm.io/) , [gorm mysql driver](https://github.com/go-gorm/mysql)

```shell
go get -u github.com/go-sql-driver/mysql
go get -u gorm.io/driver/mysql
go get -u gorm.io/gorm
```

-  [echo web framework](https://github.com/labstack/echo)

```shell
github.com/labstack/echo/v4
```

-  [go ffmpeg package](https://github.com/u2takey/ffmpeg-go)

```shell
github.com/u2takey/ffmpeg-go
```

-  [short id](https://github.com/teris-io/shortid)

```shell
github.com/teris-io/shortid
```

<br>

# 설정

### `MariaDB`

| **Variable name** | **Type** |  **Default**  |
| :---------------: | :------: | :-----------: |
|    `USERNAME`     | `string` |   `"root"`    |
|    `PASSWORD`     | `string` |   `"1111"`    |
|      `HOST`       | `string` | `"localhost"` |
|      `PORT`       | `string` |   `"3306"`    |
|     `SCHEMA`      | `string` |  `"videdit"`  |

### `Video`

| **Variable name** | **Type** | **Default**  |
| :---------------: | :------: | :----------: |
| `UPLOADFILEPATH`  | `string` | `"./upload"` |
| `OUTPUTFILEPATH`  | `string` | `"./output"` |
