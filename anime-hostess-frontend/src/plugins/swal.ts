import Swal, {SweetAlertResult} from 'sweetalert2'

export interface ISwal {
    success: (text: string) => Promise<SweetAlertResult>,
    error: (text: string) => Promise<SweetAlertResult>
}

export let swal: ISwal = {
    success(text: string) {
        return Swal.fire({
            title: '操作成功',
            text,
            icon: 'success'
        })
    },
    error(text: string) {
        return Swal.fire({
            title: '产生了错误',
            text,
            icon: 'error'
        })
    }
}
