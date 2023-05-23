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

    (_airbyte_data ->> 'id')::int as id,
    _airbyte_data ->> 'name' as name,
    _airbyte_data ->> 'email' as email,
    _airbyte_emitted_at as sync_date

from {{source('public','_airbyte_raw_users')}}

{% if is_incremental() %}
    where _airbyte_emitted_at > select max(sync_date) from {{ this }}
{% endif %}