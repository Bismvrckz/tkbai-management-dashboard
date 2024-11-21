let e = $("#kt_table_users").DataTable();
let modal = $("#modal_add_student")
let modalForm = $("#modal_add_student_form")
let newModal = new bootstrap.Modal(modal);


$("#searchUser").on("keyup", function (t) {
    e.search(t.target.value).draw()
})

$(".deleteStudent").on("click", function (t) {
    let id = $(this).attr("student-id");
})

$("#closeAddStudentModal").on("click", function (t) {
    Swal.fire({
        text: "Cancel add student?",
        icon: "warning",
        showCancelButton: !0,
        buttonsStyling: !1,
        confirmButtonText: "Yes",
        cancelButtonText: "No",
        customClass: {confirmButton: "btn btn-primary", cancelButton: "btn btn-active-light"}
    }).then((function (t) {
        t.value ? (modalForm.trigger("reset"), newModal.hide()) : ""
    }))
})


$("#discardAddStudentModal").on("click", function (t) {
    Swal.fire({
        text: "Discard student data?",
        icon: "warning",
        showCancelButton: !0,
        buttonsStyling: !1,
        confirmButtonText: "Yes",
        cancelButtonText: "No",
        customClass: {confirmButton: "btn btn-primary", cancelButton: "btn btn-active-light"}
    }).then((function (t) {
        t.value ? (modalForm.trigger("reset"), newModal.hide()) : ""
    }))
})

$(document).ready(function () {
    $(".hapus-link").on("click", function (e) {
        e.preventDefault(); // Prevent the default behavior of the anchor tag
        $(this).closest("form").submit(); // Find the closest form and submit it
    });
});

history.pushState("","","/tkbai/admin/dashboard")