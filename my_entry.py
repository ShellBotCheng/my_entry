import pymysql
import pypinyin
import hashlib
import random
import string
from zlib import crc32
from mimesis import Person

person = Person('zh')


def pinyin(word):
    s = ''
    for i in pypinyin.pinyin(word, style=pypinyin.NORMAL):
        s += ''.join(i)
    return s


def getCrc32(filename):
    return crc32(filename.encode("utf-8")) % 128


def main():
    # 打开数据库连接
    db = pymysql.connect(host='localhost',
                         user='root',
                         password='123qwe',
                         database='entry_task')

    # 使用cursor()方法获取操作游标
    cursor = db.cursor()
    # SQL 插入语句
    in_all = 0
    while in_all <= 10000000:
        nickname = person.name()
        username = pinyin(nickname) + ''.join(random.sample([str(i) for i in range(0, 9)], 8))
        tableName = getCrc32(username)
        sql = "INSERT INTO user_" + str(tableName).zfill(4) + "(username,nickname,password,salt) VALUES "
        salt = ''.join(random.sample(string.ascii_letters + string.digits, 6))
        pwd = '123456' + salt
        md_pwd = hashlib.md5(pwd.encode(encoding='utf-8')).hexdigest()
        sql = sql + '("' + username + '","' + nickname + '","' + md_pwd + '","' + salt + '");'
        in_all = in_all + 1
        cursor.execute(sql)
        db.commit()
        if in_all % 10000 == 0:
            print(in_all)


if __name__ == '__main__':
    main()
