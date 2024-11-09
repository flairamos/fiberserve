CREATE TABLE IF NOT EXISTS store (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(11) NOT NULL,
    name VARCHAR(20) NOT NULL,
    sex BOOLEAN NOT NULL,
    work INT NOT NULL,
    openid VARCHAR(36) NOT NULL,
    authentication BOOLEAN NOT NULL,
    authentication_name VARCHAR(20) NOT NULL,
    authentication_etc TEXT NOT NULL,
    background_img TEXT NOT NULL,
    following  INT NOT NULL,
    follower  INT NOT NULL,
    photos TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);


COMMENT ON TABLE store_goods IS '账户表'

COMMENT ON COLUMN posting.id IS '主键';
COMMENT ON COLUMN posting.phone IS '手机号';
COMMENT ON COLUMN posting.name IS '昵称';
COMMENT ON COLUMN posting.sex IS '性别';
COMMENT ON COLUMN posting.work IS '运动步数';
COMMENT ON COLUMN posting.openid IS 'wechat_miniprogram';
COMMENT ON COLUMN posting.authentication IS '认证状态';
COMMENT ON COLUMN posting.authentication_name IS '认证信息';
COMMENT ON COLUMN posting.authentication_etc IS '认证材料';
COMMENT ON COLUMN posting.following IS '关注';
COMMENT ON COLUMN posting.follower IS '粉丝';
COMMENT ON COLUMN posting.background_img IS '背景图片';
COMMENT ON COLUMN posting.created_at IS '创建时间';
COMMENT ON COLUMN posting.updated_at IS '更新时间';
COMMENT ON COLUMN posting.deleted_at IS '删除时间';