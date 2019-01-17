# -*- coding:utf-8 -*-
import urllib
import urllib.request
import re

page = 1
url = 'http://www.qiushibaike.com/hot/page/' + str(page)
user_agent = 'Mozilla/4.0 (compatible; MSIE 5.5; Windows NT)'
headers = { 'User-Agent' : user_agent }
request = urllib.request.Request(url,headers = headers)
response = urllib.request.urlopen(request)
content = response.read().decode('utf-8')
# print(content)
pattern = re.compile('<div class="article.*?author.*?alt="(.*?)">.*?<div class="content">.*?<span>(.*?)</span>.*?</div>.*?<div class="stats.*?class="number">(.*?)</i>',re.S)
print(pattern)
print(re.findall(pattern,content))
items = re.findall(pattern,content)
for item in re.findall(pattern,content):
    print (item[0],item[1],item[2])