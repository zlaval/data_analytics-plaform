{{
    config(
        materialized='view',
    )
}}

select
    min(name),
    count(1)
from {{ref('products')}} p
         inner join {{ref('orders')}} o on p.id = o.product_id
group by p.id
order by 1
