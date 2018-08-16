"""
Hooks for Dredd tests of API
https://dredd.readthedocs.io/en/latest/hooks-python.html
"""
import json
import dredd_hooks as hooks

UUID_EXAMPLE = "bf3ba75b-8dfe-4619-b832-31c4a087a589"
GET_ONE_INDIVIDUAL_IMPLEMENTED = False
GET_INDIVIDUALS_BY_VARIANT_IMPLEMENTED = False
GET_VARIANTS_BY_INDIVIDUAL_IMPLEMENTED = False
GET_ONE_VARIANT_IMPLEMENTED = False
GET_ONE_CALL_IMPLEMENTED = False
GET_CALLS_IMPLEMENTED = False

response_stash = {}


@hooks.after("/individuals > Get all individuals > 200 > application/json")
def save_individuals_response(transaction):
    parsed_body = json.loads(transaction['real']['body'])
    ids = [item['id'] for item in parsed_body]
    response_stash['individual_ids'] = ids


@hooks.after("/variants > Get all variants within genomic range > 200 > application/json") # noqa501
def save_variants_response(transaction):
    parsed_body = json.loads(transaction['real']['body'])
    ids = [item['id'] for item in parsed_body]
    response_stash['variant_ids'] = ids


@hooks.after("/calls > Get all calls > 200 > application/json")
def save_calls_response(transaction):
    if not GET_CALLS_IMPLEMENTED:
        transaction['skip'] = True
    else:
        parsed_body = json.loads(transaction['real']['body'])
        ids = [item['id'] for item in parsed_body]
        response_stash['call_ids'] = ids


@hooks.before("/individuals/{individual_id}/variants > Get variants called in an individual > 200 > application/json") # noqa501
def variants_by_individual(transaction):
    insert_individual_id(transaction, GET_VARIANTS_BY_INDIVIDUAL_IMPLEMENTED)

@hooks.before("/individuals/{individual_id} > Get specific individual > 200 > application/json") # noqa501
def specific_individual(transaction):
    insert_individual_id(transaction, GET_ONE_INDIVIDUAL_IMPLEMENTED)


def insert_individual_id(transaction, implemented=True):
    if not implemented:
        transaction['skip'] = True
    else:
        transaction['fullPath'] = transaction['fullPath'].replace(UUID_EXAMPLE, response_stash['individual_ids'][0]) # noqa501


@hooks.before("/calls > Add a call to the database > 201 > application/json")
def print_transaction(transaction):
    request_body = json.loads(transaction['request']['body'])
    request_body['individual_id'] = response_stash['individual_ids'][0]
    request_body['variant_id'] = response_stash['variant_ids'][0]
    transaction['request']['body'] = json.dumps(request_body)


@hooks.before("/variants/{variant_id}/individuals > Get individuals with a given variant called > 200 > application/json") # noqa501
def specific_variant(transaction):
    insert_variant_id(transaction, GET_INDIVIDUALS_BY_VARIANT_IMPLEMENTED)

@hooks.before("/variants/{variant_id} > Get specific variant > 200 > application/json") # noqa501
def individuals_by_variant(transaction):
    insert_variant_id(transaction, GET_ONE_VARIANT_IMPLEMENTED)


def insert_variant_id(transaction, implemented=True):
    if not implemented:
        transaction['skip'] = True
    else:
        transaction['fullPath'] = transaction['fullPath'].replace(UUID_EXAMPLE, response_stash['variant_ids'][0]) # noqa501


@hooks.before("/calls/{call_id} > Get specific call > 200 > application/json")
def insert_call_id(transaction):
    if not GET_ONE_CALL_IMPLEMENTED:
        transaction['skip'] = True
    else:
        transaction['fullPath'] = transaction['fullPath'].replace(UUID_EXAMPLE, response_stash['call_ids'][0]) # noqa501


@hooks.before("/individuals/{individual_id} > Get specific individual > 404 > application/json") # noqa501
@hooks.before("/variants/{variant_id} > Get specific variant > 404 > application/json") # noqa501
@hooks.before("/calls/{call_id} > Get specific call > 404 > application/json")
@hooks.before("/individuals/{individual_id}/variants > Get variants called in an individual > 404 > application/json") # noqa501
@hooks.before("/variants/{variant_id}/individuals > Get individuals with a given variant called > 404 > application/json") # noqa501
def let_pass(transaction):
    transaction['skip'] = False
