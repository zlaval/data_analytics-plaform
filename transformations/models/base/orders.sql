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

    _airbyte_data ->> '_id' as id,
    (_airbyte_data ->> 'user_id')::int as user_id,
    (_airbyte_data ->> 'product_id')::int as product_id,
    (_airbyte_data ->> 'timestamp')::timestamp as ordered_at,
    _airbyte_data ->  'details' ->> 'data' as comment,
    _airbyte_emitted_at as sync_date

from {{source('public','_airbyte_raw_orders')}}

{% if is_incremental() %}
    where _airbyte_emitted_at > select max(sync_date) from {{ this }}
{% endif %}