package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/cipepser/httpclient/sdk"
	"github.com/pkg/errors"
)

const (
	URL     = "https://poloniex.com/tradingApi"
	timeout = 10 // [sec]

	FormatSecond = "2006-01-02T15:04:05"
)

// Client wraps sdk.Client to access poloniex api.
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Logger     *log.Logger
}

// NewClient is a constructor of Client.
func NewClient(urlStr string, logger *log.Logger) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	c := &Client{
		URL:        parsedURL,
		HTTPClient: &http.Client{},
		Logger:     logger,
	}

	return c, err
}

// NewRequest is a wrapper of http.NewRequest which has timeout by context package.
func (c *Client) NewRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	return req, nil
}

// MyTradeHistory is a struct represents the history of your trades
// for each currency pair.
type MyTradeHistory []struct {
	GlobalTradeID int    `json:"globalTradeID"`
	TradeID       string `json:"tradeID"`
	Date          string `json:"date"`
	Rate          string `json:"rate"`
	Amount        string `json:"amount"`
	Total         string `json:"total"`
	Fee           string `json:"fee"`
	OrderNumber   string `json:"orderNumber"`
	Type          string `json:"type"`
	Category      string `json:"category"`
}

// MyTradeHistorys represents the history of all your trades.
type MyTradeHistorys struct {
	USDTBTC  MyTradeHistory `json:"USDT_BTC"`
	BTCETH   MyTradeHistory `json:"BTC_ETH"`
	BTCLTC   MyTradeHistory `json:"BTC_LTC"`
	USDTETH  MyTradeHistory `json:"USDT_ETH"`
	BTCBCH   MyTradeHistory `json:"BTC_BCH"`
	USDTLTC  MyTradeHistory `json:"USDT_LTC"`
	USDTBCH  MyTradeHistory `json:"USDT_BCH"`
	BTCXRP   MyTradeHistory `json:"BTC_XRP"`
	BTCETC   MyTradeHistory `json:"BTC_ETC"`
	BTCXMR   MyTradeHistory `json:"BTC_XMR"`
	BTCBTS   MyTradeHistory `json:"BTC_BTS"`
	BTCLSK   MyTradeHistory `json:"BTC_LSK"`
	USDTETC  MyTradeHistory `json:"USDT_ETC"`
	USDTXRP  MyTradeHistory `json:"USDT_XRP"`
	BTCDASH  MyTradeHistory `json:"BTC_DASH"`
	BTCZEC   MyTradeHistory `json:"BTC_ZEC"`
	BTCXEM   MyTradeHistory `json:"BTC_XEM"`
	USDTXMR  MyTradeHistory `json:"USDT_XMR"`
	USDTZEC  MyTradeHistory `json:"USDT_ZEC"`
	BTCMAID  MyTradeHistory `json:"BTC_MAID"`
	BTCSYS   MyTradeHistory `json:"BTC_SYS"`
	BTCGNT   MyTradeHistory `json:"BTC_GNT"`
	BTCFCT   MyTradeHistory `json:"BTC_FCT"`
	BTCSTR   MyTradeHistory `json:"BTC_STR"`
	BTCDGB   MyTradeHistory `json:"BTC_DGB"`
	BTCFLO   MyTradeHistory `json:"BTC_FLO"`
	BTCDOGE  MyTradeHistory `json:"BTC_DOGE"`
	USDTNXT  MyTradeHistory `json:"USDT_NXT"`
	BTCZRX   MyTradeHistory `json:"BTC_ZRX"`
	BTCSTRAT MyTradeHistory `json:"BTC_STRAT"`
	BTCNXC   MyTradeHistory `json:"BTC_NXC"`
	BTCGAME  MyTradeHistory `json:"BTC_GAME"`
	BTCSC    MyTradeHistory `json:"BTC_SC"`
	USDTDASH MyTradeHistory `json:"USDT_DASH"`
	BTCNAV   MyTradeHistory `json:"BTC_NAV"`
	BTCNXT   MyTradeHistory `json:"BTC_NXT"`
	BTCCVC   MyTradeHistory `json:"BTC_CVC"`
	BTCBURST MyTradeHistory `json:"BTC_BURST"`
	BTCDCR   MyTradeHistory `json:"BTC_DCR"`
	USDTREP  MyTradeHistory `json:"USDT_REP"`
	BTCEMC2  MyTradeHistory `json:"BTC_EMC2"`
	BTCBCN   MyTradeHistory `json:"BTC_BCN"`
	BTCARDR  MyTradeHistory `json:"BTC_ARDR"`
	USDTSTR  MyTradeHistory `json:"USDT_STR"`
	BTCSTEEM MyTradeHistory `json:"BTC_STEEM"`
	ETHZRX   MyTradeHistory `json:"ETH_ZRX"`
	BTCREP   MyTradeHistory `json:"BTC_REP"`
	BTCXPM   MyTradeHistory `json:"BTC_XPM"`
	BTCVTC   MyTradeHistory `json:"BTC_VTC"`
	BTCBTCD  MyTradeHistory `json:"BTC_BTCD"`
	BTCEXP   MyTradeHistory `json:"BTC_EXP"`
	BTCLBC   MyTradeHistory `json:"BTC_LBC"`
	ETHBCH   MyTradeHistory `json:"ETH_BCH"`
	ETHETC   MyTradeHistory `json:"ETH_ETC"`
	ETHZEC   MyTradeHistory `json:"ETH_ZEC"`
	BTCCLAM  MyTradeHistory `json:"BTC_CLAM"`
	ETHLSK   MyTradeHistory `json:"ETH_LSK"`
	BTCPOT   MyTradeHistory `json:"BTC_POT"`
	BTCFLDC  MyTradeHistory `json:"BTC_FLDC"`
	BTCPINK  MyTradeHistory `json:"BTC_PINK"`
	ETHGNT   MyTradeHistory `json:"ETH_GNT"`
	BTCBLK   MyTradeHistory `json:"BTC_BLK"`
	BTCXCP   MyTradeHistory `json:"BTC_XCP"`
	BTCNOTE  MyTradeHistory `json:"BTC_NOTE"`
	BTCAMP   MyTradeHistory `json:"BTC_AMP"`
	BTCVRC   MyTradeHistory `json:"BTC_VRC"`
	BTCGNO   MyTradeHistory `json:"BTC_GNO"`
	BTCRIC   MyTradeHistory `json:"BTC_RIC"`
	BTCXBC   MyTradeHistory `json:"BTC_XBC"`
	BTCBCY   MyTradeHistory `json:"BTC_BCY"`
	BTCVIA   MyTradeHistory `json:"BTC_VIA"`
	BTCOMNI  MyTradeHistory `json:"BTC_OMNI"`
	XMRBLK   MyTradeHistory `json:"XMR_BLK"`
	BTCRADS  MyTradeHistory `json:"BTC_RADS"`
	BTCSJCX  MyTradeHistory `json:"BTC_SJCX"`
	XMRDASH  MyTradeHistory `json:"XMR_DASH"`
	BTCBELA  MyTradeHistory `json:"BTC_BELA"`
	BTCPPC   MyTradeHistory `json:"BTC_PPC"`
	BTCNEOS  MyTradeHistory `json:"BTC_NEOS"`
	BTCNMC   MyTradeHistory `json:"BTC_NMC"`
	BTCSBD   MyTradeHistory `json:"BTC_SBD"`
	BTCPASC  MyTradeHistory `json:"BTC_PASC"`
	ETHREP   MyTradeHistory `json:"ETH_REP"`
	ETHCVC   MyTradeHistory `json:"ETH_CVC"`
	XMRLTC   MyTradeHistory `json:"XMR_LTC"`
	BTCGRC   MyTradeHistory `json:"BTC_GRC"`
	BTCXVC   MyTradeHistory `json:"BTC_XVC"`
	BTCBTM   MyTradeHistory `json:"BTC_BTM"`
	BTCNAUT  MyTradeHistory `json:"BTC_NAUT"`
	ETHSTEEM MyTradeHistory `json:"ETH_STEEM"`
	ETHGNO   MyTradeHistory `json:"ETH_GNO"`
	BTCHUC   MyTradeHistory `json:"BTC_HUC"`
	XMRBTCD  MyTradeHistory `json:"XMR_BTCD"`
	XMRMAID  MyTradeHistory `json:"XMR_MAID"`
	XMRZEC   MyTradeHistory `json:"XMR_ZEC"`
	XMRNXT   MyTradeHistory `json:"XMR_NXT"`
	XMRBCN   MyTradeHistory `json:"XMR_BCN"`
}

// GetMyTradeHistory get your trade history by poloniex api.
// [PARAMTERS]
// start, end: specify the range of the history you wanna get.
// These parameters, start and end, are formated as "yyyyyyyy"
// If you wish to get ALL your trade history, you specify the parameters as below.
// "start" should be a time before your first trade.
// "end" should be a time after your last order or now.
func (c *Client) GetMyTradeHistory(start, end string) (MyTradeHistorys, error) {
	s, err := time.Parse(FormatSecond, start)
	if err != nil {
		return MyTradeHistorys{}, err
	}
	e, err := time.Parse(FormatSecond, end)
	if err != nil {
		return MyTradeHistorys{}, err
	}

	fmt.Println(strconv.FormatInt(s.Unix(), 10))
	fmt.Println(strconv.FormatInt(e.Unix(), 10))

	// set timeout timer by context package.
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	// construct a query
	command := "returnTradeHistory"
	nonce := strconv.Itoa(int(time.Now().Unix()))

	query := url.Values{}
	query.Set("command", command)
	query.Set("nonce", nonce)
	query.Set("currencyPair", "all")
	query.Set("start", strconv.FormatInt(s.Unix(), 10))
	query.Set("end", strconv.FormatInt(e.Unix(), 10))
	body := query.Encode()

	req, err := c.NewRequest(ctx, "POST", "", strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	// set authentication header to req
	setPrivateHeader(req, body)

	// send a http request and get a response.
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return MyTradeHistorys{}, err
	}

	// decode http response as type MyTradeHistory.
	h := MyTradeHistorys{}
	err = sdk.DecodeBody(resp, &h)
	if err != nil {
		return MyTradeHistorys{}, err
	}

	return h, nil
}

// setPrivateHeader sets authentication header to req.
func setPrivateHeader(req *http.Request, body string) {
	key := os.Getenv("POLOKEY")
	secret := os.Getenv("POLOSECRET")

	sign := makeHMAC(body, secret)

	req.Header.Set("Sign", sign)
	req.Header.Set("Key", key)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

// makeHMAC returns a HMAC by sha512.
func makeHMAC(msg, key string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}

func main() {
	c, _ := NewClient(URL, nil)

	start := "2017-05-18T15:00:00"
	end := "2017-05-19T15:00:00"

	h, err := c.GetMyTradeHistory(start, end)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h)
}
