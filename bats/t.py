'''
@Description: In User Settings Edit
@Author: your name
@Date: 2019-08-20 20:51:24
@LastEditTime: 2019-08-21 11:27:10
@LastEditors: Please set LastEditors
'''
# -*- coding: utf-8 -*-
import sys
from datetime import datetime
import time
import random
#reload(sys)
#sys.setdefaultencoding('utf8')

def formatLog(str):
    t = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
    print("[{}]{}".format(t, str))

class TM_producs(object):
    def __init__(self,storename, cookie, title):
        self.storename = storename
        self.url = 'https://{}.m.tmall.com'.format(storename)
        self.headers = {
                #":authority": "juehuai.m.tmall.com",
                #":method": "GET",
                #":path": "//shop/shop_auction_search.do?sort=s&p=1&page_size=12&from=h5&ajson=1&_tm_source=tmallsearch&callback=jsonp_83739921",
                #":scheme": "https",
                "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
                "accept-encoding": "gzip, deflate, br",
                "accept-language": "zh-CN,zh;q=0.9",
                "cache-control": "no-cache",
                "cookie": cookie,
                "pragma": "no-cache",
                "upgrade-insecure-requests": "1",
                "user-agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
        }
        datenum = datetime.now().strftime('%Y%m%d%H%M')
        self.filename = '{}_{}.csv'.format(self.storename, datenum)
        self.title = title


    def main(self):
        '''循环爬取所有页面宝贝'''
        total_page = 3
        for i in range(1,total_page+1):
            formatLog('总计{}页商品，已经提取第{}页'.format(total_page,i))
            time.sleep(1+random.random())

if __name__ == '__main__':
    formatLog("start...")
    storename = 'juehuai'
    if len(sys.argv) < 3:
        formatLog("params error")
        sys.exit()
    storename = sys.argv[1]
    cookie = sys.argv[2]
    formatLog("start {}".format(storename))
    if len(sys.argv) == 4:
        title = sys.argv[3].split(",")
    else:
        title = ['item_id','price','quantity','sold','title','totalSoldQuantity','url','img']
    tm = TM_producs(storename, cookie, title)
    tm.main()
    formatLog("success")
    formatLog("end {}".format(storename))
