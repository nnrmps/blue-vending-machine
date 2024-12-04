import { Button, ConfigProvider } from 'antd';
import { createStyles } from 'antd-style';
import { ButtonType, ButtonVariantType } from 'antd/es/button';

const useStyle = createStyles(({ prefixCls, css }) => ({
  linearGradientButton: css`
    &.${prefixCls}-btn-primary:not([disabled]):not(
        .${prefixCls}-btn-dangerous
      ) {
      border-width: 0;

      > span {
        position: relative;
      }

      &::before {
        content: '';
        background: linear-gradient(135deg, #6253e1, #04befe);
        position: absolute;
        inset: 0;
        opacity: 1;
        transition: all 0.3s;
        border-radius: inherit;
      }

      &:hover::before {
        opacity: 0;
      }
    }
  `,
}));

type ButtonStyleProps = {
  icon?: React.ReactNode;
  iconPosition?: 'start' | 'end';
  onClick?: () => void;
  title?: string;
  type?: ButtonType;
  variant?: ButtonVariantType;
  disabled?: boolean;
};

export const GradientButton = ({
  icon,
  iconPosition = 'start',
  onClick,
  title,
  disabled = false,
  type = 'primary',
  variant = 'solid',
}: ButtonStyleProps) => {
  const { styles } = useStyle();

  return (
    <ConfigProvider
      button={{
        className: styles.linearGradientButton,
      }}
    >
      <Button
        type={type}
        disabled={disabled}
        variant={variant}
        icon={icon}
        iconPosition={iconPosition}
        onClick={onClick}
      >
        {title}
      </Button>
    </ConfigProvider>
  );
};
