-- Modify "users" table
ALTER TABLE "users" ADD COLUMN "status" character varying NOT NULL DEFAULT 'A', ADD COLUMN "name" character varying NOT NULL, ADD COLUMN "username" character varying NOT NULL, ADD COLUMN "password" character varying NOT NULL, ADD COLUMN "email" character varying NOT NULL, ADD COLUMN "date_joined" character varying NULL, ADD COLUMN "timezone" character varying NOT NULL DEFAULT 'TimezoneAmericaLosAngeles', ADD COLUMN "profile_pic_url" character varying NULL, ADD COLUMN "thumbnail_url" character varying NULL, ADD COLUMN "phone_number" character varying NULL, ADD COLUMN "is_admin" boolean NOT NULL DEFAULT false, ADD COLUMN "is_super_admin" boolean NOT NULL DEFAULT false, ADD COLUMN "business_unit_id" uuid NOT NULL, ADD COLUMN "organization_id" uuid NOT NULL, ADD CONSTRAINT "users_business_units_business_unit" FOREIGN KEY ("business_unit_id") REFERENCES "business_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "users_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;