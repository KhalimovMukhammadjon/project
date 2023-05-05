SELECT
    o.order_id,
    oi.product_id,
    oi.quantity,
    oi.list_price,
    oi.discount,
    c.order_limit_price,
    c.name,
    c.discount
FROM orders AS o
LEFT JOIN order_item AS oi ON oi.order_id=o.order_id
LEFT JOIN 


SELECT 
    order_limit_price,
    discount
FROM 
    coupons
WHERE name =$1


SELECT
    list_price * quantity AS sum

FROM 
    order_item
WHERE 
    order_id = $1