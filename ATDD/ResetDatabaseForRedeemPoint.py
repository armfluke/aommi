#!/usr/bin/python
# coding=utf-8
import MySQLdb

db = MySQLdb.connect(host="178.128.48.140", user="root", passwd="Admin123!", db="aommi")

# you must create a Cursor object. It will let
#  you execute all the queries you need
cursor = db.cursor()

# Use all the SQL you like
# cursor.execute("SELECT ")
cursor.execute("UPDATE account SET PointBalance=1000 WHERE AccountID='1969900224728';")
cursor.execute("UPDATE account SET PointBalance=100 WHERE AccountID='1101500889111';")
cursor.execute("UPDATE account SET PointBalance=10 WHERE AccountID='1103702074462';")
db.commit()

# print all the first cell of all the rows
# for row in cursor.fetchall():
#     print row[:]

db.close()