const body = $("body");
const dataSchemaInputTemplate = Handlebars.compile($("#data-schema-input-template").html());

let fieldId = 0;

body.on("click", ".data-schema-input-data-type-select", function (e) {
    let inputId = $(this).data("id");

    const selected = $(this).val();

    if (selected === "array") {
        fieldId = 0
        let elementsSchema = $(`#data-schema-input-elements-schema-${inputId}`)
        elementsSchema.show()
        elementsSchema.html(dataSchemaInputTemplate({id: `${inputId}-elements`, title: "Array Elements Schema"}))
        $(`#data-schema-input-fields-schema-${inputId}`).hide()
    } else if (selected === "object") {
        fieldId += 1
        let fieldsSchema = $(`#data-schema-input-fields-schema-${inputId}`)
        fieldsSchema.show()
        let fieldsSchemaContainer = $(`#data-schema-input-fields-schema-container-${inputId}`)
        fieldsSchemaContainer.html(dataSchemaInputTemplate({id: `${inputId}-fields-${fieldId}`, title: `Object Field Schema`, field: true}))
        $(`#data-schema-input-elements-schema-${inputId}`).hide()
    } else {
        fieldId = 0
        $(`#data-schema-input-elements-schema-${inputId}`).hide()
        $(`#data-schema-input-fields-schema-${inputId}`).hide()
    }
})

body.on("click", ".data-schema-input-fields-schema-add-btn", function (e) {
    e.preventDefault()
    let inputId = $(this).data("id")
    fieldId += 1
    let fieldsSchemaContainer = $(`#data-schema-input-fields-schema-container-${inputId}`)
    fieldsSchemaContainer.append(dataSchemaInputTemplate({id: `${inputId}-fields-${fieldId}`, title: `Object Field Schema`, field: true}))
})

body.on("click", ".data-schema-input-delete-field-btn", function (e) {
    e.preventDefault()
    let inputId = $(this).data("id")
    $(`#data-schema-input-${inputId}`).remove()
})