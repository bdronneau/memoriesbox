ALTER TABLE "mbox"."memories" DROP CONSTRAINT "pk_memories_id";
ALTER TABLE "mbox"."memories" ADD PRIMARY KEY ("xid");
ALTER TABLE "mbox"."memories" DROP COLUMN "id";
