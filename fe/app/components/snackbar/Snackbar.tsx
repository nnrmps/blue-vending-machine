import Icon, {
  CheckCircleOutlined,
  CloseCircleOutlined,
} from '@ant-design/icons';
import * as React from 'react';
import { tv, type VariantProps } from 'tailwind-variants';

export type SnackbarVariant = 'error' | 'success';

const snackbar = tv({
  slots: {
    base: 'caption-4 flex w-[480px] -translate-y-6 scale-x-75 transform items-center justify-between rounded-[16px] px-[24px] py-[16px] text-[#ffffff] opacity-0 shadow-md transition-all duration-300',
    wrapper: 'flex flex-1 space-x-[8px]',
    actionWrapper: 'flex items-center space-x-[16px]',
  },
  variants: {
    variant: {
      success: 'bg-[#0d995c]',
      error: 'bg-[#db1439]',
    },
    visible: {
      true: 'translate-y-0 scale-x-100 opacity-100',
    },
    center: {
      true: {
        wrapper: 'justify-center',
      },
    },
    origin: {
      component: {
        base: 'fixed left-[50%] top-[64px] z-[1000] translate-x-[-50%]',
      },
      hook: {},
    },
  },
});

type SnackbarVariants = VariantProps<typeof snackbar>;

export interface SnackbarProps extends SnackbarVariants {
  content: React.ReactNode;
  disallowClose?: boolean;
  action?: React.ReactNode;
  onClose?: () => void;
}
export const Snackbar = ({
  content,
  origin = 'component',
  variant = 'success',
  visible = false,
  center = false,
}: SnackbarProps) => {
  const { base, wrapper } = snackbar({
    variant,
    visible,
    center,
    origin,
  });

  return (
    <div className={base()}>
      <div className={wrapper()}>
        {variant === 'success' ? (
          <CheckCircleOutlined />
        ) : (
          <CloseCircleOutlined />
        )}
        <div
          data-testid={content}
          className='whitespace-pre-wrap typography-caption-4'
        >
          {content}
        </div>
      </div>
    </div>
  );
};
