CREATE TABLE "professionalUser" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamp DEFAULT (now()) NOT NULL,
  "name" varchar NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "email" varchar NOT NULL,
  "date_of_birth" Date NOT NULL,
  "cpf" integer NOT NULL,
  "image_id" bigint NOT NULL,
  "updated_at" timestamp DEFAULT (now()) NOT NULL,
  "subjectMatter_id" integer UNIQUE NOT NULL,
  "subjectMatter_class_id" integer UNIQUE NOT NULL,
  "class_hour_price" varchar NOT NULL,
  "calendar_id" integer UNIQUE NOT NULL
);

CREATE TABLE "calendar" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "subjectMatter_id" integer UNIQUE NOT NULL,
  "time" timestamp NOT NULL,
  "date" Date NOT NULL,
  "available" boolean NOT NULL,
  "filled_student_id" integer UNIQUE NOT NULL,
  "professionalUser_id" integer UNIQUE NOT NULL
  
);

CREATE TABLE "subjectMatter" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()) NOT NULL,
  "category" varchar NOT NULL,
  "abstract" varchar NOT NULL
);

CREATE TABLE "subjectMatterClass" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamp DEFAULT (now()) NOT NULL,
  "subjectMatter_id" integer UNIQUE NOT NULL,
  "professional_id" integer NOT NULL,
  "durantion" integer NOT NULL,
  "enrollment_date" Date NOT NULL,
  "enrollment_time" timestamp NOT NULL,
  "cancellation" boolean NOT NULL,
  "cancellation_reason" varchar NOT NULL,
  "student_attendence" boolean NOT NULL,
  "study_material" varchar NOT NULL,
  "testing_exam" varchar NOT NULL
);

CREATE TABLE "studentUser" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "date_of_birth" Date NOT NULL,
  "gender" varchar NOT NULL,
  "created_at" timestamp 
   DEFAULT (now()) NOT NULL,
  "responsible_student_id" integer UNIQUE NOT NULL,
  "updated_at" timestamp 
   DEFAULT (now()) NOT NULL,
  "subjectMatter_class_id" integer UNIQUE NOT NULL,
  "calendar_id" integer UNIQUE NOT NULL
);

CREATE TABLE "responsibleStudent" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "email" varchar NOT NULL,
  "date_of_birth" Date NOT NULL,
  "cpf" integer NOT NULL,
  "phone_id" bigint UNIQUE NOT NULL,
  "created_at" timestamp
   DEFAULT (now()) NOT NULL,
  "updated_at" timestamp 
   DEFAULT (now()) NOT NULL
);

CREATE TABLE "professionalInformation" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "experience_period" varchar NOT NULL,
  "ocupation_area" varchar NOT NULL,
  "university" varchar NOT NULL,
  "graduation_diploma" varchar NOT NULL,
  "validate" boolean NOT NULL,
  "graduation_country" varchar NOT NULL,
  "graduation_city" varchar NOT NULL,
  "graduation_state" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()) NOT NULL,
  "updated_at" timestamp DEFAULT (now()) NOT NULL,
  "professional_user_id" bigint UNIQUE NOT NULL
);

CREATE TABLE "phone" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "country_code" integer NOT NULL,
  "area_core" integer NOT NULL,
  "number" integer NOT NULL,
  "type" varchar NOT NULL,
  "created_at" timestamp 
   DEFAULT (now()) NOT NULL,
  "updated_at" timestamp 
   DEFAULT (now()) NOT NULL,
  "user_id" bigint UNIQUE NOT NULL
);

ALTER TABLE "phone" ADD FOREIGN KEY ("user_id") REFERENCES "professionalUser" ("id");

ALTER TABLE "professionalInformation" ADD FOREIGN KEY ("professional_user_id") REFERENCES "professionalUser" ("id");

ALTER TABLE  "subjectMatterClass" ADD FOREIGN KEY ("id") REFERENCES "professionalUser" ("subjectMatter_class_id");

ALTER TABLE  "subjectMatter" ADD FOREIGN KEY ("id") REFERENCES "professionalUser"  ("subjectMatter_id");

ALTER TABLE "calendar" ADD FOREIGN KEY ("id") REFERENCES "professionalUser" ("calendar_id");

ALTER TABLE "calendar" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");

ALTER TABLE "calendar" ADD FOREIGN KEY ("filled_student_id") REFERENCES "studentUser" ("id");

ALTER TABLE "subjectMatterClass" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");

ALTER TABLE "studentUser" ADD FOREIGN KEY ("responsible_student_id") REFERENCES "responsibleStudent" ("id");

ALTER TABLE "studentUser" ADD FOREIGN KEY ("subjectMatter_class_id") REFERENCES "subjectMatterClass"  ("id");

ALTER TABLE "studentUser" ADD FOREIGN KEY ("calendar_id") REFERENCES "calendar" ("id");

ALTER TABLE "responsibleStudent" ADD FOREIGN KEY ("phone_id") REFERENCES "phone" ("id");

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
