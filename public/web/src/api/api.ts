import type { AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import axios from "axios";

const service: AxiosInstance = axios.create({
    timeout: 10 * 1000, // 请求超时时间
    headers: { "Content-Type": "application/json;charset=UTF-8" }
});
service.interceptors.request.use((config: AxiosRequestConfig) => {
    // 这里可以设置token: config!.headers!.Authorization = token
    return config;
});
service.interceptors.response.use((response: AxiosResponse) => {
    const data = response.data;
    if (data.code === 200) {
        return data;
    } else {
        return Promise.reject(data);
    }
},
(err) => {
    return Promise.reject(err.response);
}
);
const request = {
    get<T = any>(url: string, data?: any): Promise<T> {
        return request.request("GET", url, { params: data });
    },
    post<T = any>(url: string, data?: any): Promise<T> {
        return request.request("POST", url, { data });
    },
    put<T = any>(url: string, data?: any): Promise<T> {
        return request.request("PUT", url, { data });
    },
    delete<T = any>(url: string, data?: any): Promise<T> {
        return request.request("DELETE", url, { params: data });
    },
    request<T = any>(method = "GET", url: string, data?: any): Promise<T> {
        return new Promise((resolve, reject) => {
            service({ method, url, ...data })
                .then((res) => {
                    resolve(res as unknown as Promise<T>);
                })
                .catch((e: Error | AxiosError) => {
                    reject(e);
                })
        });
    }
};

export default request;