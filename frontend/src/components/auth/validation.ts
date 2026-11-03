import * as Yup from 'yup';

export const USERNAME_REGEX = /^[a-zA-Z0-9_]+$/;
export const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]{2,}$/;
export const PRINTABLE_ASCII_REGEX = /^[\x20-\x7E]+$/;

export const USERNAME_MIN = 4;
export const USERNAME_MAX = 15;
export const PASSWORD_MIN = 8;
export const PASSWORD_MAX = 64;

const emailField = Yup.string()
    .trim()
    .required('Введите email')
    .matches(EMAIL_REGEX, 'Некорректный email');

export const signinSchema = Yup.object().shape({
    email: Yup.string().trim().required('Введите email'),
    password: Yup.string().required('Введите пароль'),
});

export const signupSchema = Yup.object().shape({
    username: Yup.string()
        .required('Введите никнейм')
        .min(USERNAME_MIN, `Никнейм: от ${USERNAME_MIN} до ${USERNAME_MAX} символов`)
        .max(USERNAME_MAX, `Никнейм: от ${USERNAME_MIN} до ${USERNAME_MAX} символов`)
        .matches(USERNAME_REGEX, 'Никнейм может содержать только a-z, A-Z, 0-9 и _'),
    email: emailField,
    password: Yup.string()
        .required('Введите пароль')
        .min(PASSWORD_MIN, `Пароль: от ${PASSWORD_MIN} до ${PASSWORD_MAX} символов`)
        .max(PASSWORD_MAX, `Пароль: от ${PASSWORD_MIN} до ${PASSWORD_MAX} символов`)
        .matches(PRINTABLE_ASCII_REGEX, 'Пароль может содержать только печатные символы ASCII'),
});

export const isFilled = (values: Record<string, string>): boolean =>
    Object.values(values).every((value) => value.trim() !== '');
