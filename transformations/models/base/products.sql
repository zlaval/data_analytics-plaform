{{
    config(
        materialized='incremental',
        unique_key='id',
        indexes=[
            {'columns': ['id'], 'type': 'btree', 'unique': true},
            {'columns': ['sync_date'], 'type': 'btree'}
        ]
    )
}}

select

    ph.product_id as id,
    COALESCE (max(ph.name),max(t.name) )as name,
    (max(array[to_json(ph.modified_at)::text,ph.price::text]))[2]::int as price,
    max(ph.sync_date) as sync_date


from
    {{ref('products_history')}} ph
    left outer join {{this}} t on t.id=ph.product_id

{% if is_incremental() %}
    where ph.id in(
        select id from {{ref('products_history')}}
        group by id
        having max(sync_date) > (select max(sync_date) from {{ this }})
    )
{% endif %}

group by ph.product_id

