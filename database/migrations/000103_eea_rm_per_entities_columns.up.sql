-- Copyright 2024 Stacklok, Inc
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--      http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

BEGIN;

ALTER TABLE entity_execution_lock DROP COLUMN IF EXISTS repository_id;
ALTER TABLE entity_execution_lock DROP COLUMN IF EXISTS artifact_id;
ALTER TABLE entity_execution_lock DROP COLUMN IF EXISTS pull_request_id;

ALTER TABLE flush_cache DROP COLUMN IF EXISTS repository_id;
ALTER TABLE flush_cache DROP COLUMN IF EXISTS artifact_id;
ALTER TABLE flush_cache DROP COLUMN IF EXISTS pull_request_id;

-- make project_id not nullable
ALTER TABLE entity_execution_lock ALTER COLUMN project_id SET NOT NULL;
ALTER TABLE flush_cache ALTER COLUMN project_id SET NOT NULL;

COMMIT;
