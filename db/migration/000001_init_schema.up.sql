CREATE TABLE "professionalUser" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "name" varchar NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "email" varchar NOT NULL,
  "date_of_birth" Date NOT NULL,
  "cpf" integer NOT NULL,
  "image_id" bigint NOT NULL,
  "phone_id" bigint UNIQUE NOT NULL,
  "professional_information_id" bigint UNIQUE NOT NULL,
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "subjectMatter_id" integer UNIQUE NOT NULL UNIQUE,
  "subjectMatter_class_id" integer UNIQUE NOT NULL,
  "class_hour_price" varchar NOT NULL,
  "calendar_id" integer UNIQUE NOT NULL
);

CREATE TABLE "calendar" (
  "id" bigserial PRIMARY KEY,
  "subjectMatter_id" integer UNIQUE NOT NULL,
  "time" timestamp,
  "date" Date,
  "available" boolean,
  "filled_student_id" integer UNIQUE NOT NULL,
  "professionalUser_id" integer UNIQUE NOT NULL
);

CREATE TABLE "subjectMatter" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "category" varchar NOT NULL,
  "abstract" varchar
);

CREATE TABLE "subjectMatterClass" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "subjectMatter_id" integer UNIQUE NOT NULL,
  "professional_id" integer NOT NULL,
  "durantion" integer NOT NULL,
  "enrollment_date" Date NOT NULL,
  "enrollment_time" timestamp NOT NULL,
  "cancellation" boolean,
  "cancellation_reason" varchar,
  "student_attendence" boolean,
  "study_material" varchar,
  "testing_exam" varchar
);

CREATE TABLE "studentUser" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "name" varchar,
  "date_of_birth" Date NOT NULL,
  "gender" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "responsible_student_id" integer UNIQUE NOT NULL,
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "subjectMatter_class_id" integer UNIQUE NOT NULL, 
  "calendar_id" integer UNIQUE NOT NULL
);

CREATE TABLE "responsibleStudent" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "email" varchar NOT NULL,
  "date_of_birth" Date NOT NULL,
  "cpf" integer NOT NULL,
  "phone_id" bigint UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "professionalInformation" (
  "id" bigserial PRIMARY KEY,
  "experience_period" varchar,
  "ocupation_area" varchar,
  "university" varchar,
  "graduation_diploma" varchar NOT NULL,
  "validate" boolean,
  "graduation_country" varchar,
  "graduation_city" varchar,
  "graduation_state" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "phone" (
  "id" bigserial PRIMARY KEY,
  "country_code" integer NOT NULL,
  "area_core" integer NOT NULL,
  "number" integer NOT NULL,
  "type" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
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
  "professionalUser_id" bigserial UNIQUE NOT NULL,
  "subjectMatter_id" bigserial UNIQUE NOT NULL,
  PRIMARY KEY ("professionalUser_id", "subjectMatter_id")
);

ALTER TABLE "professionalUser_subjectMatter" ADD FOREIGN KEY ("professionalUser_id") REFERENCES "professionalUser" ("id");

ALTER TABLE "professionalUser_subjectMatter" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");


CREATE TABLE "subjectMatter_calendar" (
  "subjectMatter_id" bigserial UNIQUE NOT NULL, 
  "calendar_id" bigserial UNIQUE NOT NULL,
  PRIMARY KEY ("subjectMatter_id", "calendar_id")
);

ALTER TABLE "subjectMatter_calendar" ADD FOREIGN KEY ("subjectMatter_id") REFERENCES "subjectMatter" ("id");

ALTER TABLE "subjectMatter_calendar" ADD FOREIGN KEY ("calendar_id") REFERENCES "calendar" ("id");
