# Function And Parameter Naming Conventions

## Folder
### Folder Name
เป็นตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
src
```

## File
### File Name
เป็นตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
aommi.go
```

## Package
### Package Name
เป็นตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
aommi
```

## Test Function
### Function Name
รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ SNAKE เช่น
```
Test_ActionToTest_Input_152_Should_Be_150
```

### Parameter Name
1. ถ้าชื่อตัวแปรเป็นคำๆเดียวให้ตั้งชื่อเป็นพิมพ์เล็กทั้งหมด เช่น
```
day, month, year
```
2. ถ้าชื่อตัวแปรมีความยาวตั้งแต่ 2 คำขึ้นไป ให้คำหลังขึ้นตันด้วยตัวอักษรตัวใหญ่เสมอ ในรูปแบบ Camel เช่น
```
startDay, endMonth
```
3. ถ้าชื่อตัวแปรเก็บค่าที่เป็นพหูพจน์ ให้เติม s ต่อท้ายตัวแปรเสมอ เช่น
```
days, months
```
4. ชื่อตัวแปรใน struct ให้ตังชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ Camel เช่น
```
ResultDay, ResultMonth
```
5. ถ้าเป็นตัวแปร Constant ให้ตังชื่อเป็นตัวอักษรพิมพ์ใหญ่ทั้งหมด เช่น
```
HOUR, MINUTE
```

### Addtion
1. หลัง Comma ทุกครั้ง ให้เว้นวรรค 1 space
2. ชื่อ Struct ให้ตังชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์เล็กและคำถัดๆไปด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ Camel เช่น
```
durationDate
```

## Function
### Function Name
1. ถ้าเป็นฟังก์ชันที่เรียกใช้เฉพาะในแพคเกจนั้นๆให้ตั้งชื่อขึ้นต้นด้วยตัวอักษรพิมพ์เล็ก ในรูปแบบ Camel เช่น
```
func transferHour
```
2. ถ้าเป็นฟังก์ชันที่เรียกใช้จากภายนอกแพคเกจให้ตั้งชื่อขึ้นต้นด้วยตััวอักษรพิมพ์ใหญ่ ในรูปแบบ Camel เช่น
```
func TransferMinute
```

### Parameter Name
1. ถ้าชื่อตัวแปรเป็นคำๆเดียวให้ตั้งชื่อเป็นพิมพ์เล็กทั้งหมด เช่น
```
day, month, year
```
2. ถ้าชื่อตัวแปรมีความยาวตั้งแต่ 2 คำขึ้นไป ให้คำหลังขึ้นตันด้วยตัวอักษรตัวใหญ่เสมอ ในรูปแบบ Camel เช่น
```
startDay, endMonth
```
3. ถ้าชื่อตัวแปรเก็บค่าที่เป็นพหูพจน์ ให้เติม s ต่อท้ายตัวแปรเสมอ เช่น
```
days, months
```
4. ชื่อตัวแปรใน struct ให้ตังชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ Camel เช่น
```
ResultDay, ResultMonth
```
5. ถ้าเป็นตัวแปร Constant ให้ตังชื่อเป็นตัวอักษรพิมพ์ใหญ่ทั้งหมด เช่น
```
HOUR, MINUTE
```

### Addtion
1. หลัง Comma ทุกครั้ง ให้เว้นวรรค 1 space
2. ชื่อ Struct ให้ตังชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์เล็กและคำถัดๆไปด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ Camel เช่น
```
durationDate
```

### ข้อตกลง commit message ใน github
`[Created]: สร้างไฟล์ใหม่`

`[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code`

`[Added]: กรณีเพิ่ม function , function test ใหม่เข้ามา`

`[Deleted]: ลบไฟล์ออก`

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน

## คำสั่ง Run Test

### คำสั่ง Run Acceptance Test (API)
```
newman run filename
```

## Additional
[See more git and go command](https://github.com/ImKK-000/git-and-go-step)

## Version 

## Deploy on server
#Server port
3001

#Login to server
1. ssh root@178.128.48.140
2. password: lpk,=eokPdb0

#go location
/usr/local/go/bin

#Add go to path
export PATH=$PATH:/usr/local/go/bin

#project location
root/bootcamp/aommi/aommi/API

#set GOPATH
cd /bootcamp/aommi/aommi/API
export GOPATH=`pwd`

#Build and run
1. go build -o aommi src/main.go
2. ./aommi &

#Check process and kill
1. fg #check running background
2. ps aux | grep aommi #get process id
3. kill PROCESS_ID

