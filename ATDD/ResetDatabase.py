#!/usr/bin/python
# coding=utf-8
import MySQLdb

db = MySQLdb.connect(host="178.128.48.140", user="root", passwd="Admin123!", db="aommi")

cursor = db.cursor()

cursor.execute("UPDATE account SET PointBalance=500, SavingAccount=0, FixedAccount=0 WHERE AccountID='1969900224728';")
cursor.execute("UPDATE account SET PointBalance=500, SavingAccount=0, FixedAccount=0 WHERE AccountID='1101500889111';")
cursor.execute("UPDATE account SET PointBalance=500, SavingAccount=0, FixedAccount=0  WHERE AccountID='1100400758552';")
cursor.execute("UPDATE account SET PointBalance=500, SavingAccount=0, FixedAccount=0  WHERE AccountID='1103702074462';")
cursor.execute("UPDATE account SET PointBalance=500, SavingAccount=0, FixedAccount=0  WHERE AccountID='1140100074828';")

db.commit()
db.close()