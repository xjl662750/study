#导入包，读取数据
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
# %matplotlib inline

data_train = pd.read_csv(r'D:/Goxx/python/titanic/train.csv')
data_test = pd.read_csv(r'D:/Goxx/python/titanic/test.csv')
#我们将测试集导入，再将删除Survived数据的训练集与测验集进行合并，这样便于进行数据处理

y = data_train['Survived'] # 将训练集Survived 数据存储在y中
del data_train['Survived'] # 删除训练集Survived数据

sum_id = data_test['PassengerId'] # 存储测试集乘客ID

df = pd.merge(data_train, data_test,how='outer') # 合并无Survived数据的训练集与测验集，how = ‘outer’ 意为并集

#删掉无关因素
df = df.drop(['Name','PassengerId','Ticket','Cabin'],axis=1) # 删除姓名、ID、船票信息、客舱信息，axis=0 删除行，=1 删除列

#缺失数据填充
df['Age'] = df['Age'].fillna(df['Age'].mean())  # 用平均值填充空值
df['Fare'] = df['Fare'].fillna(df['Fare'].mean())
df['Embarked'] = df['Embarked'].fillna( df['Embarked'].value_counts().index[0]) # 用数量最多项填充

#将性别与港口用哑变量表示
dumm = pd.get_dummies(df[['Sex','Embarked']]) # '哑变量'矩阵
df = df.join(dumm)
del df['Sex']
del df['Embarked']

#数据降维
df['Age'] = (df['Age']-df['Age'].min()) /(df['Age'].max()-df['Age'].min())

df['Fare'] = (df['Fare']-df['Fare'].min()) /(df['Fare'].max()-df['Fare'].min())

#训练模型
data_train = df[:len(data_train)] # 将合并后的数据分离
data_test = df[len(data_train):]

from sklearn.model_selection import train_test_split

X_train, X_val, y_train,y_val = train_test_split(data_train,y,test_size=0.3, random_state=42) # 以7：3（0.3）将训练集与获救结果随机拆分，随机种子为42

from sklearn.linear_model import LogisticRegression # 引入逻辑回归
LR = LogisticRegression()

LR.fit(X_train, y_train) # 训练数据
print('训练集准确率：\n',LR.score(X_train, y_train)) # 分数
print('验证集准确率：\n',LR.score(X_val, y_val))

#预测测验集
pred= LR.predict(data_test) # pred 为预测结果

pred = pd.DataFrame({'PassengerId':sum_id.values, 'Survived':pred}) # 格式化预测结果

pred.to_csv('pred_LR.csv',index=None) # 导出数据