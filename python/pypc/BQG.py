# -*- coding:utf-8 -*-
import urllib
from urllib import parse
import urllib.request
import re
import threading
import time
import ssl
import json
 
#糗事百科爬虫类
class BQG:
 
    #初始化方法，定义一些变量
    def __init__(self):
        # self.pageIndex = 1
        self.user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'
        #初始化headers
        self.headers = { 'User-Agent' : self.user_agent,}
        #存放段子的变量，每一个元素是每一页的段子们
        self.stories = []
        #存放程序是否继续运行的变量
        self.enable = False

    # #获取书url
    # def getBook(self,bookname):
    #     try:
    #         url = 'https://www.biqudu.com/searchbook.php?keyword='+bookname
    #         #构建请求的request
    #         # print(url)
    #         request = urllib.request.Request(url,headers = self.headers)
    #         #利用urlopen获取页面代码
    #         response = urllib.request.urlopen(request)
    #         # print (response.read().decode('utf-8'))
    #         #将页面转化为UTF-8编码
    #         pageCode = response.read().decode('utf-8')
    #         # print(pageCode)
    #         if not pageCode:
    #             print ("页面加载失败....")
    #             return None
    #         pattern = re.compile('<div class="item">.*?<span>(.*?)</span>.*?<a href="(.*?)">(.*?)</a>',re.S)
    #         items = re.findall(pattern,pageCode)
    #         books = []
    #         #遍历正则表达式匹配的信息
    #         for item in items:
    #                 books.append([item[0].strip(),item[1].strip(),item[2].strip()])
    #                 #item[0]是作者，item[1]是url，item[2]是书名
    #         return books
 
    #     except urllib.request.URLError(e):
    #         if hasattr(e,"reason"):
    #             print (u"连接失败,错误原因",e.reason)
    #             return None

    #获得章节名和url
    def getPage(self,bookname):
        try:
            url = 'http://www.biquge.com.tw/modules/article/soshu.php?searchkey=+'+bookname
            #构建请求的request
            request = urllib.request.Request(url,headers = self.headers)
            #利用urlopen获取页面代码
            response = urllib.request.urlopen(request)
            #将页面转化为UTF-8编码
            pageCode = response.read().decode('gbk')
            print(pageCode)
            if not pageCode:
                print ("页面加载失败....")
                return None
            pattern = re.compile('<dd><a href="(.*?)">(.*?)</a>',re.S)
            items = re.findall(pattern,pageCode)
            chapters = []
            #遍历正则表达式匹配的信息
            for item in items:
                    chapters.append([item[0].strip(),item[1].strip()])
                    #item[0]是章节url，item[1]是章节名
            # print(chapters)
            return chapters
 
        except urllib.request.URLError(e):
            if hasattr(e,"reason"):
                print (u"连接失败,错误原因",e.reason)
                return None
 
 
    #传入章节url，返回本页信息
    def getPageItems(self,chapterurl):
        url = 'http://www.biquge.com.tw'+chapterurl
        # print(url)
         #构建请求的request
        request = urllib.request.Request(url,headers = self.headers)
        #利用urlopen获取页面代码
        response = urllib.request.urlopen(request)
         #将页面转化为UTF-8编码
        pageCode = response.read().decode('gbk')
        # pageCode = self.getPage(chapterurl)
        if not pageCode:
            print ("页面加载失败....")
            return None
        # print(pageCode)
        pattern = re.compile('<div class="bookname">.*?<h1>(.*?)</h1>.*?<div id="content">(.*?)</div>',re.S)
        # print(pattern)
        # f=open("activity_order.html","rb").read().decode("utf-8")
        # print(pageCode)
        items = re.findall(pattern,pageCode)
        # print(items)
        #用来存储每页的段子们
        pageStories = []
        #遍历正则表达式匹配的信息
        for item in items:
                replaceBR = re.compile('<br />')
                text1 = re.sub(replaceBR,"",item[1])
                replaceBR1 = re.compile('&nbsp;')
                text = re.sub(replaceBR1,"",text1)
                # replaceBR2 = re.compile('\n')
                # text = re.sub(replaceBR2,"",text2)
                #item[0]是一个段子的发布者，item[1]是内容，item[2]是是点赞数
                # print(len(text))
                pageStories.append([item[0].strip(),text.strip()])
        return pageStories
 
    #加载并提取页面的内容，加入到列表中
    # def loadPage(self):
    #     #如果当前未看的页数少于2页，则加载新一页
    #     if self.enable == True:
    #         if len(self.stories) < 2:
    #             #获取新一页
    #             pageStories = self.getPageItems()
    #             #将该页的段子存放到全局list中
    #             if pageStories:
    #                 self.stories.append(pageStories)
    #                 #获取完之后页码索引加一，表示下次读取下一页
    #                 # self.pageIndex += 1
    
    # #调用该方法，每次敲回车打印输出一个段子
    # def getOneStory(self,pageStories):
    #     #遍历一页的段子
    #     for story in pageStories:
    #         #等待用户输入
    #         inputt = input()
    #         #每当输入回车一次，判断一下是否要加载新页面
    #         self.loadPage()
    #         #如果输入Q则程序结束
    #         if inputt == "Q":
    #             self.enable = False
    #             return
    #         print (u"章节名:%s\t内容:%s\n%s" %(story[0],story[1]))
    
    #开始方法
    def start(self):
        print (u"请输入书名")
        #使变量为True，程序可以正常运行
        # self.enable = True
        #先加载一页内容
        # self.loadPage()
        #局部变量，控制当前读到了第几页
        # nowPage = 0
        # while self.enable:
        #     if len(self.stories)>0:
        #         #从全局list中获取一页的段子
        #         pageStories = self.stories[0]
        #         #当前读到的页数加一
        #         # nowPage += 1
        #         #将全局list中第一个元素删除，因为已经取出
        #         del self.stories[0]
        #         #输出该页的段子
        #         self.getOneStory(pageStories)
        ssl._create_default_https_context = ssl._create_unverified_context
        bookname1 = input()
        bookname=parse.quote(bookname1.encode("gbk"))
        print(bookname)
        chapters=self.getPage(bookname)
        print(chapters)
        result=[]
        for chapter in chapters:
            pageStories=self.getPageItems(chapter[0])
            for story in pageStories:
                result.append([story[0].strip(),story[1].strip()])
                # print (u"章节名:%s\t内容:%s\n" %(story[0],story[1]))

        with open('result.json', 'w', encoding='UTF-8') as file:
            file.write(json.dumps(result, indent=4, ensure_ascii=False))







        # pageStories=self.getPageItems()
        # for story in pageStories:
        #     print (u"章节名:%s\t内容:%s\n" %(story[0],story[1]))
 
 
spider = BQG()
spider.start()