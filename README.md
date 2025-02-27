## ใน Project นี้เป็นโปรเจคที่ใช้ GORM เป็น structure โปรแกรมขนาดเล็ก โดยใช้ GORM, Gin, Swagger, Viper และ goose
- structure โปรเจคคร่าวๆ
  - ในโปรเจคจะมี folder ที่แยกไว้โดยหน้าที่แต่ละ folder จะมีดังนี้
  - config ใช้สำหรับ config database driver และ อื่นๆ
  - db ใช้สำหรับจัดการกับ database และการ migration
  - docs ส่วนนี้จะ genarate ยังการใช้ swagger
  - dto (Data Transfer Object) เป็น folder ที่จะคอยเก็บ object ต่างๆ เช่น request เวลารับ request เราจำเป็นต้องมี Object หรือ Structure ในการรับค่าของ request ไว้เพราะบางครั้ง request ที่ส่งมาอาจจะมีข้อมูลเพิ่มเติมจาก Model เดิมจึงสร้างไว้เพื่อความยืดหยุ่น
  - internal ใน folder นี้จะมีการเก็บ controller service repository ต่างๆไว้ และ Model ของข้อมูล
    - Controller ทำหน้าที่รับ Request และ ส่ง Response ให้แก่ฝั่ง Client
    - Service เป็น Buisness Layer เวลาเราจะจัดการข้อมูลที่ได้รับมา จะตรวจสอบ หรือกรองข้อมูล กระกระทำทั้งหมดจะอยู่ใน service นี้ แล้วส่งไปที่ repository แต่ถ้าฝั่ง repository มีการ return ข้อมูลกลับมาก็จะนำข้อมูลที่ได้รับมาส่งกลับไปที่ controller
    - Repository เป็นในส่วนการจัดการกับ Database หลักจากรับ request มาแล้วก็จะนำ request ไปทำการ query ต่างๆ เอาง่ายๆ เป็นส่วนในการติดต่อฐานข้อมูล แล้วส่งกลับไปที่ service
    - Model ทำหน้าที่เป็นตัวเก็บ structure หรือ Object ของข้อมูลจาก table ของฐานข้อมูล

- วิธีสร้างไฟล์ migration sql
- goose -dir ./db/migration create create_products sql
- goose -dir ./db/migration create create_users sql
- โดยในไฟล์ migration จะเป็นการสร้าง table และ กำหนด field ต่างๆ

![image](https://github.com/user-attachments/assets/763e9769-f4b3-4ffe-8268-8877b4a3360b)
- คำสั่งรัน migration : goose -dir ./db/migration sqlite3 coffee.db up

- run swagger ครั้งแรกใช้ swag init --parseDependency --parseInternal แต่ถ้าเคยรันแล้วมีการเพิ่ม action เพิ่มเช่น เพิ่มการลบข้อมูลมาให้ใช้แค่ swag init

# สรุปสั้นๆ
- ใช้ GORM กับ SQLite
- ใช้ Gin ทำ API
- ใช้ Swagger สำหรับ API Docs
- ใช้ Goose ทำ Migration
- แยกโค้ดตาม Clean Architecture

