CREATE TABLE "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE "individuals" (
"id" TEXT PRIMARY KEY,
"description" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "variants" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"chromosome" TEXT NOT NULL,
"start" integer NOT NULL,
"ref" TEXT NOT NULL,
"alt" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE "calls" (
"id" TEXT PRIMARY KEY,
"individual" char(36) NOT NULL,
"variant" char(36) NOT NULL,
"genotype" TEXT NOT NULL,
"format" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
