CREATE TABLE IF NOT EXISTS store_goods (
    id SERIAL PRIMARY KEY,
    store_id INT NOT NULL,
    goods_name VARCHAR(20) NOT NULL,
    content VARCHAR(1000) NOT NULL,
    photos TEXT NOT NULL,
    price NUMERIC(5,2),
    status BOOLEAN NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMENT ON TABLE store_goods IS '商铺商品'

COMMENT ON COLUMN posting.id IS '主键';
COMMENT ON COLUMN posting.store_id IS '商铺id';
COMMENT ON COLUMN posting.goods_name IS '商品名称';
COMMENT ON COLUMN posting.content IS '商品介绍';
COMMENT ON COLUMN posting.photos IS '商品图片';
COMMENT ON COLUMN posting.price IS '商品价格';
COMMENT ON COLUMN posting.status IS '上下架状态';
COMMENT ON COLUMN posting.created_at IS '创建时间';
COMMENT ON COLUMN posting.updated_at IS '更新时间';
COMMENT ON COLUMN posting.deleted_at IS '删除时间';