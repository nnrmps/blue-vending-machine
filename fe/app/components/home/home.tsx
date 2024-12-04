import { useMemo, useState } from 'react';
import { Button, Col, Row, Statistic } from 'antd';
import { UndoOutlined } from '@ant-design/icons';
import { Screen } from '../screen/screen';
export const Home = () => {
  const [coins1, setCoins1] = useState(0);
  const [coins5, setCoins5] = useState(0);
  const [coins10, setCoins10] = useState(0);
  const [bank20, setBank20] = useState(0);
  const [bank50, setBank50] = useState(0);
  const [bank100, setBank100] = useState(0);
  const [bank500, setBank500] = useState(0);
  const [bank1000, setBank1000] = useState(0);

  const sumPrice = useMemo(() => {
    return (
      coins1 * 1 +
      coins5 * 5 +
      coins10 * 10 +
      bank20 * 20 +
      bank50 * 50 +
      bank100 * 100 +
      bank500 * 500 +
      bank1000 * 1000
    );
  }, [coins1, coins5, coins10, bank20, bank50, bank100, bank500, bank1000]);

  const totalPrice = {
    coins1,
    coins5,
    coins10,
    bank20,
    bank50,
    bank100,
    bank500,
    bank1000,
  };

  const handleResetMoney = () => {
    setCoins1(0);
    setCoins5(0);
    setCoins10(0);
    setBank20(0);
    setBank50(0);
    setBank100(0);
    setBank500(0);
    setBank1000(0);
  };

  return (
    <>
      <Row gutter={[16, 24]} style={{ alignItems: 'center' }}>
        <Col span={24}>
          <div className='flex gap-[24px] justify-between items-center'>
            <Statistic
              title='Account Balance (THB)'
              value={sumPrice}
              precision={2}
            />

            <Button
              style={{ float: 'right' }}
              type='default'
              danger
              onClick={handleResetMoney}
              icon={<UndoOutlined />}
              iconPosition='start'
              shape='round'
              disabled={sumPrice === 0}
            >
              Refund
            </Button>
          </div>
        </Col>

        <Col className='gutter-row' span={24}>
          <Screen
            totalPrice={totalPrice}
            sumPrice={sumPrice}
            handleClearData={handleResetMoney}
          />
        </Col>

        <Col className='gutter-row' span={6}>
          <Button
            style={{ minHeight: '50px' }}
            block
            onClick={() => setCoins1((prev) => prev + 1)}
          >
            1
          </Button>
        </Col>
        <Col className='gutter-row' span={6}>
          <Button
            style={{ minHeight: '50px' }}
            block
            onClick={() => setCoins5((prev) => prev + 1)}
          >
            5
          </Button>
        </Col>
        <Col className='gutter-row' span={6}>
          <Button
            style={{ minHeight: '50px' }}
            block
            onClick={() => setCoins10((prev) => prev + 1)}
          >
            10
          </Button>
        </Col>
        <Col className='gutter-row' span={6}>
          <Button
            style={{ minHeight: '50px' }}
            block
            onClick={() => setBank20((prev) => prev + 1)}
          >
            20
          </Button>
        </Col>
        <Col className='gutter-row' span={6}>
          <Button
            style={{ minHeight: '50px' }}
            block
            onClick={() => setBank50((prev) => prev + 1)}
          >
            50
          </Button>
        </Col>
        <Col className='gutter-row' span={6}>
          <Button
            style={{ minHeight: '50px' }}
            block
            onClick={() => setBank100((prev) => prev + 1)}
          >
            100
          </Button>
        </Col>
        <Col className='gutter-row' span={6}>
          <Button
            style={{ minHeight: '50px' }}
            block
            onClick={() => setBank500((prev) => prev + 1)}
          >
            500
          </Button>
        </Col>
        <Col className='gutter-row' span={6}>
          <Button
            style={{ minHeight: '50px' }}
            block
            onClick={() => setBank1000((prev) => prev + 1)}
          >
            1000
          </Button>
        </Col>
      </Row>
    </>
  );
};
