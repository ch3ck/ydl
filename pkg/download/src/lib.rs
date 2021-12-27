//! -*- mode: rust; -*-
//!
//! download - downloads youtube files
extern crate libc;
use env_logger;
use libc::c_char;
use log::{debug, info};
use rustube;
use std::ffi::CStr;

#[no_mangle]
pub async extern "C" fn download(
    c_url: *const c_char,
    c_path: *const c_char,
) -> Result<(), Box<dyn std::error::Error>> {
    env_logger::init();

    // convert str in C to rust safely
    let pre_url = unsafe {
        assert!(!c_url.is_null());
        CStr::from_ptr(c_url)
    };

    let url = pre_url.to_str().unwrap();
    url.chars().count() as u32;

    // convert str for path from C to rust safely
    let pre_path = unsafe {
        assert!(!c_path.is_null());
        CStr::from_ptr(c_path)
    };
    let path = pre_path.to_str().unwrap();
    path.chars().count() as u32;

    // download video
    info!("video_url: {:?}", url);
    let _result = rustube::download_best_quality(url).await?;
    debug!("download status: {:?}", _result);

    Ok(())
}

#[cfg(test)]
pub mod tests {
    use super::*;
    use libc;
    use std::ffi::CString;

    #[tokio::test]
    async fn test_download() {
        let url = CString::new("https://www.youtube.com/watch?v=lWEbEtr_Vng")
            .unwrap();
        url.as_ptr() as *const libc::c_char;

        let fp = CString::new("~/Downloads").unwrap();
        fp.as_ptr() as *const libc::c_char;

        download(
            url.as_ptr() as *const libc::c_char,
            fp.as_ptr() as *const libc::c_char,
        )
        .await
        .unwrap();
    }
}
