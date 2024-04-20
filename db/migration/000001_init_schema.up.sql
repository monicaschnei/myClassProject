CREATE TABLE "professionalUser" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "created_at" timestamp DEFAULT (now()) NOT NULL,
  "name" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "gender" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "date_of_birth" Date NOT NULL,
  "cpf" varchar UNIQUE NOT NULL,
  "image_id" bigint UNIQUE NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT (now()),
  "hashed_password" varchar NOT NULL,
  "updated_at" timestamp DEFAULT (now()) NOT NULL,
  "class_hour_price" varchar NOT NULL
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
  "subjectMatter_id" integer NOT NULL,
  "durantion" integer NOT NULL,
  "enrollment_date" Date NOT NULL,
  "enrollment_time" timestamp NOT NULL,
  "cancellation" boolean NOT NULL,
  "cancellation_reason" varchar NOT NULL,
  "student_attendence" boolean NOT NULL,
  "study_material" varchar NOT NULL,
  "testing_exam" varchar NOT NULL,
  "professional_user_id" bigint NOT NULL,
  "student_user_id" bigint NOT NULL
);

CREATE TABLE "studentUser" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "date_of_birth" Date NOT NULL,
  "gender" varchar NOT NULL,
  "created_at" timestamp
   DEFAULT (now()) NOT NULL,
  "responsible_student_id" bigserial NOT NULL,
  "updated_at" timestamp
   DEFAULT (now()) NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT (now()),
  "hashed_password" varchar NOT NULL
);

CREATE TABLE "responsibleStudent" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "email" varchar NOT NULL,
  "date_of_birth" Date NOT NULL,
  "cpf" varchar UNIQUE NOT NULL,
  "created_at" timestamp
   DEFAULT (now()) NOT NULL,
  "updated_at" timestamp 
   DEFAULT (now()) NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT (now()),
  "hashed_password" varchar NOT NULL
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
  "professional_user_id" bigint NOT NULL
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

CREATE TABLE "availability" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "date" varchar NOT NULL,
    "start" varchar NOT NULL,
    "end_time" varchar NOT NULL,
    "is_available" boolean NOT NULL,
    "user_id" bigint UNIQUE NOT NULL,
    "username" varchar UNIQUE NOT NULL
);

ALTER TABLE "professionalInformation" ADD FOREIGN KEY ("professional_user_id") REFERENCES "professionalUser" ("id");

ALTER TABLE "subjectMatterClass" ADD FOREIGN KEY ("professional_user_id") REFERENCES "professionalUser" ("id");

ALTER TABLE "subjectMatterClass" ADD FOREIGN KEY ("student_user_id") REFERENCES "studentUser" ("id");

ALTER TABLE "subjectMatterClass" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");

ALTER TABLE "studentUser" ADD FOREIGN KEY ("responsible_student_id") REFERENCES "responsibleStudent" ("id");

CREATE TABLE "professionalUser_subjectMatter" (
  "professionalUser_id" bigserial
  ,
  "subjectMatter_id" bigserial
  ,
  PRIMARY KEY ("professionalUser_id", "subjectMatter_id")
);

ALTER TABLE "professionalUser_subjectMatter" ADD FOREIGN KEY ("professionalUser_id") REFERENCES "professionalUser" ("id");

ALTER TABLE "professionalUser_subjectMatter" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");

