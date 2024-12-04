import * as React from 'react';
import hotToast, { ToastOptions } from 'react-hot-toast';
import { Snackbar, SnackbarProps } from './Snackbar';

const DEFAULT_DURATION = 3000;

interface ToastProps extends SnackbarProps {
  id?: string;
  autoClose?: boolean;
  duration?: number;
}

export const useToast = () => {
  const handleClose = (props: ToastProps, toastId?: string) => {
    return props.disallowClose
      ? undefined
      : () => {
          props.onClose?.();
          hotToast.dismiss(toastId);
        };
  };

  const toastOption = ({ autoClose = true, id }: ToastProps): ToastOptions => {
    return {
      duration: autoClose ? DEFAULT_DURATION : Infinity,
      position: 'top-center',
      id,
    };
  };

  const toast = {
    success: (props: ToastProps) =>
      hotToast.custom(
        (t) => (
          <Snackbar
            variant='success'
            origin='hook'
            {...t}
            {...props}
            onClose={handleClose(props, t.id)}
          />
        ),
        toastOption(props)
      ),
    error: (props: ToastProps) =>
      hotToast.custom(
        (t) => (
          <Snackbar
            variant='error'
            origin='hook'
            {...t}
            {...props}
            onClose={handleClose(props, t.id)}
          />
        ),
        toastOption(props)
      ),
    close: (toastId?: string) => hotToast.dismiss(toastId),
  };

  return {
    toast,
  };
};
