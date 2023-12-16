CREATE TABLE "professionalUser" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp  DEFAULT (now()),
  "name" varchar ,
  "username" varchar ,
  "password" varchar ,
  "gender" varchar ,
  "email" varchar ,
  "date_of_birth" Date ,
  "cpf" integer ,
  "image_id" bigint ,
  "phone_id" bigint UNIQUE ,
  "professional_information_id" bigint UNIQUE ,
  "updated_at" timestamp  DEFAULT (now()),
  "subjectMatter_id" integer UNIQUE ,
  "subjectMatter_class_id" integer UNIQUE ,
  "class_hour_price" varchar ,
  "calendar_id" integer UNIQUE 
);

CREATE TABLE "calendar" (
  "id" bigserial PRIMARY KEY,
  "subjectMatter_id" integer UNIQUE 
  ,
  "time" timestamp,
  "date" Date,
  "available" boolean,
  "filled_student_id" integer UNIQUE 
  ,
  "professionalUser_id" integer UNIQUE 
  
);

CREATE TABLE "subjectMatter" (
  "id" bigserial PRIMARY KEY,
  "title" varchar 
  ,
  "created_at" timestamp 
   DEFAULT (now()),
  "category" varchar 
  ,
  "abstract" varchar
);

CREATE TABLE "subjectMatterClass" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp 
   DEFAULT (now()),
  "subjectMatter_id" integer UNIQUE 
  ,
  "professional_id" integer 
  ,
  "durantion" integer 
  ,
  "enrollment_date" Date 
  ,
  "enrollment_time" timestamp 
  ,
  "cancellation" boolean,
  "cancellation_reason" varchar,
  "student_attendence" boolean,
  "study_material" varchar,
  "testing_exam" varchar
);

CREATE TABLE "studentUser" (
  "id" bigserial PRIMARY KEY,
  "username" varchar 
  ,
  "password" varchar 
  ,
  "name" varchar,
  "date_of_birth" Date 
  ,
  "gender" varchar,
  "created_at" timestamp 
   DEFAULT (now()),
  "responsible_student_id" integer UNIQUE 
  ,
  "updated_at" timestamp 
   DEFAULT (now()),
  "subjectMatter_class_id" integer UNIQUE 
  , 
  "calendar_id" integer UNIQUE 
  
);

CREATE TABLE "responsibleStudent" (
  "id" bigserial PRIMARY KEY,
  "name" varchar 
  ,
  "gender" varchar 
  ,
  "email" varchar 
  ,
  "date_of_birth" Date 
  ,
  "cpf" integer 
  ,
  "phone_id" bigint UNIQUE 
  ,
  "created_at" timestamp 
   DEFAULT (now()),
  "updated_at" timestamp 
   DEFAULT (now())
);

CREATE TABLE "professionalInformation" (
  "id" bigserial PRIMARY KEY,
  "experience_period" varchar 
  ,
  "ocupation_area" varchar 
  ,
  "university" varchar 
  ,
  "graduation_diploma" varchar 
  ,
  "validate" boolean 
  ,
  "graduation_country" varchar 
  ,
  "graduation_city" varchar 
  ,
  "graduation_state" varchar 
  ,
  "created_at" timestamp 
   DEFAULT (now()),
  "updated_at" timestamp 
   DEFAULT (now())
);

CREATE TABLE "phone" (
  "id" bigserial PRIMARY KEY,
  "country_code" integer 
  ,
  "area_core" integer 
  ,
  "number" integer 
  ,
  "type" varchar 
  ,
  "created_at" timestamp 
   DEFAULT (now()),
  "updated_at" timestamp 
   DEFAULT (now())
);

ALTER TABLE "phone" ADD FOREIGN KEY ("id") REFERENCES "professionalUser" ("phone_id");

ALTER TABLE "professionalInformation" ADD FOREIGN KEY ("id") REFERENCES "professionalUser" ("professional_information_id");

ALTER TABLE "subjectMatterClass" ADD FOREIGN KEY ("id") REFERENCES "professionalUser" ("subjectMatter_class_id");

ALTER TABLE "professionalUser" ADD FOREIGN KEY ("calendar_id") REFERENCES "calendar" ("id");

ALTER TABLE "calendar" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");

ALTER TABLE "calendar" ADD FOREIGN KEY ("filled_student_id") REFERENCES "studentUser" ("id");

ALTER TABLE "subjectMatterClass" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");

ALTER TABLE "studentUser" ADD FOREIGN KEY ("responsible_student_id") REFERENCES "responsibleStudent" ("id");

ALTER TABLE "subjectMatterClass" ADD FOREIGN KEY ("id") REFERENCES "studentUser" ("subjectMatter_class_id");

ALTER TABLE "studentUser" ADD FOREIGN KEY ("calendar_id") REFERENCES "calendar" ("id");

ALTER TABLE "phone" ADD FOREIGN KEY ("id") REFERENCES "responsibleStudent" ("phone_id");

CREATE TABLE "professionalUser_subjectMatter" (
  "professionalUser_id" bigserial UNIQUE 
  ,
  "subjectMatter_id" bigserial UNIQUE 
  ,
  PRIMARY KEY ("professionalUser_id", "subjectMatter_id")
);

ALTER TABLE "professionalUser_subjectMatter" ADD FOREIGN KEY ("professionalUser_id") REFERENCES "professionalUser" ("id");

ALTER TABLE "professionalUser_subjectMatter" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");


CREATE TABLE "subjectMatter_calendar" (
  "subjectMatter_id" bigserial UNIQUE 
  , 
  "calendar_id" bigserial UNIQUE 
  ,
  PRIMARY KEY ("subjectMatter_id", "calendar_id")
);

ALTER TABLE "subjectMatter_calendar" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");

ALTER TABLE "subjectMatter_calendar" ADD FOREIGN KEY ("calendar_id") REFERENCES "calendar" ("id");
