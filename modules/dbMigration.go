package modules

// 创建数据表
const schemaUsers = `
CREATE TABLE IF NOT EXISTS public.users
(
    user_id uuid NOT NULL,
    user_name character varying(150) NOT NULL,
    user_first_name character varying(150),
    user_last_name character varying(150),
    user_email character varying(255),
    user_token character varying(40) NOT NULL,
    user_password character varying(128),
    PRIMARY KEY (user_id)
);
`

const schemaProjects = `
CREATE TABLE IF NOT EXISTS public.projects
(
    table_id serial NOT NULL,
    project_id integer NOT NULL,
    PRIMARY KEY (table_id),
    CONSTRAINT "pid唯一确定一个项目" UNIQUE (project_id)
);
`

const schemaReleases = `
CREATE TABLE IF NOT EXISTS public.releases
(
    table_id serial NOT NULL,
    release_version character varying(200) NOT NULL,
    last_commit_hash character varying(1000) NOT NULL,
    project_id integer NOT NULL,
    PRIMARY KEY (table_id),
    CONSTRAINT "pid和version唯一确定一个release" UNIQUE (release_version, project_id),
    CONSTRAINT "对应pid" FOREIGN KEY (project_id)
        REFERENCES public.projects (project_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);
`

const schemaCommits = `
CREATE TABLE IF NOT EXISTS public.commits
(
    table_id bigserial NOT NULL,
    hash character varying(1000) NOT NULL,
    "time" character varying(1000) NOT NULL,
    author character varying(1000) NOT NULL,
    email character varying(1000) NOT NULL,
    release_table_id integer NOT NULL,
    PRIMARY KEY (table_id),
    CONSTRAINT "hash和release唯一确定一行" UNIQUE (hash, release_table_id),
    CONSTRAINT "对应的release" FOREIGN KEY (release_table_id)
        REFERENCES public.releases (table_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);
`

const schemaObjects = `
CREATE TABLE IF NOT EXISTS public.objects
(
    table_id bigserial NOT NULL,
    parameters character varying(10000) NOT NULL,
    hash character varying(1000) NOT NULL,
    start_line integer NOT NULL DEFAULT 0,
    end_line integer NOT NULL DEFAULT 0,
    object_path character varying(1000) NOT NULL,
    current_object_id character varying(1000) NOT NULL,
    father_object_id character varying(1000),
    old_line integer NOT NULL DEFAULT 0,
    new_line integer NOT NULL DEFAULT 0,
    deleted_line integer NOT NULL DEFAULT 0,
    added_line integer NOT NULL DEFAULT 0,
    release_table_id integer NOT NULL,
    commit_table_id bigint NOT NULL,
    PRIMARY KEY (table_id),
    CONSTRAINT "唯一确定object亦称函数" UNIQUE (parameters, hash, start_line, object_path, current_object_id, commit_table_id, end_line),
    CONSTRAINT "唯一确定release" FOREIGN KEY (release_table_id)
        REFERENCES public.releases (table_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT "唯一确定commit" FOREIGN KEY (commit_table_id)
        REFERENCES public.commits (table_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);
`

const schemaNodes = `
CREATE TABLE IF NOT EXISTS public.nodes
(
    table_id bigserial NOT NULL,
    object_path character varying(1000) NOT NULL,
    object_parameters character varying(10000) NOT NULL DEFAULT '',
    current_object_id character varying(1000) NOT NULL,
    father_object_id character varying(1000) NOT NULL,
    old_confidence double precision NOT NULL DEFAULT 0.0,
    new_confidence double precision NOT NULL DEFAULT 0.0,
    commit_table_id bigint NOT NULL,
    object_table_id bigint NOT NULL,
    PRIMARY KEY (table_id),
    UNIQUE (object_path, object_parameters, current_object_id),
    CONSTRAINT "方便返回commitHistory" FOREIGN KEY (commit_table_id)
        REFERENCES public.commits (table_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT "方便返回objectHistory" FOREIGN KEY (object_table_id)
        REFERENCES public.objects (table_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
);
`

const indexObjectPath = "create index path_idx on nodes(object_path);"

func Create() {
	_, err = Db.Raw(schemaUsers).Rows() // 执行原生SQL语句建表
	_, err = Db.Raw(schemaProjects).Rows()
	_, err = Db.Raw(schemaReleases).Rows()
	_, err = Db.Raw(schemaCommits).Rows()
	_, err = Db.Raw(schemaObjects).Rows()
	_, err = Db.Raw(schemaNodes).Rows()
	_, err = Db.Raw(indexObjectPath).Rows()
	if err != nil {
		err.Error()
	}
}
