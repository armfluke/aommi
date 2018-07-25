#!/usr/bin/python
# coding=utf-8
import MySQLdb

db = MySQLdb.connect(host="178.128.48.140", user="root", passwd="Admin123!", db="aommi")

# you must create a Cursor object. It will let
#  you execute all the queries you need
cursor = db.cursor()

# Use all the SQL you like
# cursor.execute("SELECT ")
cursor.execute("UPDATE account SET PointBalance=0, SavingAccount=0 WHERE AccountID='1969900224728';")
cursor.execute("UPDATE account SET PointBalance=10, FixedAccount=0 WHERE AccountID='1101500889111';")
cursor.execute("UPDATE account SET PointBalance=300, SavingAccount=0, FixedAccount=0  WHERE AccountID='1100400758552';")
cursor.execute("UPDATE account SET PointBalance=150, SavingAccount=0, FixedAccount=0  WHERE AccountID='1103702074462';")
cursor.execute("UPDATE account SET PointBalance=450, SavingAccount=0, FixedAccount=0  WHERE AccountID='1140100074828';")

cursor.execute("DELETE FROM codeUsed WHERE AccountID='1100400758552' AND Date(DateEarn)=CURDATE() AND CodeID = 3")
cursor.execute("INSERT INTO codeUsed (AccountID,CodeID,PointEarn) VALUES ('1100400758552',3,20)")

cursor.execute("DELETE FROM codeUsed WHERE AccountID='1103702074462' AND Date(DateEarn)=CURDATE() AND CodeID = 3")
cursor.execute("INSERT INTO codeUsed (AccountID,CodeID,PointEarn) VALUES ('1103702074462',3,20),('1103702074462',3,20)")

cursor.execute("DELETE FROM codeUsed WHERE AccountID='1140100074828' AND Date(DateEarn)=CURDATE() AND CodeID = 3")
cursor.execute("INSERT INTO codeUsed (AccountID,CodeID,PointEarn) VALUES ('1100400758552',3,20),('1100400758552',3,20),('1100400758552',3,20),('1100400758552',3,20)")

db.commit()

# print all the first cell of all the rows
# for row in cursor.fetchall():
#     print row[:]

db.close()