async function checkID() {
    try {
        let certificate_id = document.getElementById("certificate-id");
        let inputResponse = document.getElementById("certificate-id-input-response");
        let certificate_holder = document.getElementById("certificate-holder");
        let certInputResponse = document.getElementById("certificate-holder-input-response");

        if (!certificate_id.value || !certificate_holder.value) {
            inputResponse.innerHTML = "<div>Tolong masukan ID anda</div>";
            certInputResponse.innerHTML = "<div>Tolong masukan Nama Pemegang</div>";

            return Swal.fire({
                icon: "error",
                title: "Gagal",
                text: "Tolong isi field dengan benar!",
                buttonsStyling: false,
                confirmButtonText: "Ok",
                customClass: {
                    confirmButton: "btn btn-danger",
                },
            });
        }

        const url =
            `${API_URL}/certificate/validate/id/${certificate_id.value}/name/${certificate_holder.value}`;

        const validate = await fetch(url);

        const resValidate = await validate.json();

        if (validate.status === 200) {
            Swal.fire({
                text: "Sertifikat anda valid!",
                icon: "success",
                buttonsStyling: false,
                confirmButtonText: "Selanjutnya",
                customClass: {
                    confirmButton: "btn btn-primary",
                },
            }).then((result) => {
                // if (result.isConfirmed) {
                //     window.location.href = BASE_URL + `/certificate/${resValidate.additionalInfo.testID}/name/${resValidate.additionalInfo.name}`;
                // }
            });
        } else {
            return Swal.fire({
                icon: "error",
                title: "Gagal",
                text: resValidate.message,
            });
        }
    } catch (error) {
        return Swal.fire({
            icon: "error",
            title: "Login Gagal",
            text: error,
        });
    }
}
