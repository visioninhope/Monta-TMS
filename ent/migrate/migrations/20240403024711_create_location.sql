-- Create "locations" table
CREATE TABLE "locations" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "version" bigint NOT NULL DEFAULT 1, "status" character varying NOT NULL DEFAULT 'A', "code" character varying NOT NULL, "name" character varying NOT NULL, "description" text NULL, "address_line_1" character varying(150) NOT NULL, "address_line_2" character varying(150) NULL, "city" character varying(150) NOT NULL, "postal_code" character varying(10) NOT NULL, "longitude" double precision NULL, "latitude" double precision NULL, "place_id" character varying(255) NULL, "is_geocoded" boolean NOT NULL DEFAULT false, "business_unit_id" uuid NOT NULL, "organization_id" uuid NOT NULL, "location_category_id" uuid NULL, "state_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "locations_business_units_business_unit" FOREIGN KEY ("business_unit_id") REFERENCES "business_units" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "locations_location_categories_location_category" FOREIGN KEY ("location_category_id") REFERENCES "location_categories" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "locations_organizations_organization" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "locations_us_states_state" FOREIGN KEY ("state_id") REFERENCES "us_states" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "location_code_organization_id" to table: "locations"
CREATE UNIQUE INDEX "location_code_organization_id" ON "locations" ("code", "organization_id");
