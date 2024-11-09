CREATE TABLE IF NOT EXISTS posting (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(11) NOT NULL,
    name VARCHAR(20) NOT NULL,
    authentication BOOLEAN NOT NULL,
    authentication_name VARCHAR(20) NOT NULL,
    type SMALLINT NOT NULL,
    like_people_phone VARCHAR(11) NOT NULL,
    like_people_name VARCHAR(20) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    photos TEXT NOT NULL,
    like_number BIGINT NOT NULL,
    comment_number BIGINT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE notice IS '帖子'

COMMENT ON COLUMN posting.id IS '主键';
COMMENT ON COLUMN posting.phone IS '手机号';
COMMENT ON COLUMN posting.name IS '昵称';
COMMENT ON COLUMN posting.authentication IS '是否认证';
COMMENT ON COLUMN posting.authentication_name IS '认证信息';
COMMENT ON COLUMN posting.type IS '帖子类型';
COMMENT ON COLUMN posting.like_people_phone IS '被展现人手机号';
COMMENT ON COLUMN posting.like_people_name IS '被展现人昵称';
COMMENT ON COLUMN posting.content IS '帖子内容';
COMMENT ON COLUMN posting.photos IS '帖子图片';
COMMENT ON COLUMN posting.like_number IS '点赞数';
COMMENT ON COLUMN posting.comment_number IS '评论数';
COMMENT ON COLUMN posting.created_at IS '创建时间';
COMMENT ON COLUMN posting.updated_at IS '更新时间';
COMMENT ON COLUMN posting.deleted_at IS '删除时间';