import { Table, Collapse } from 'antd';
import PropTypes from 'prop-types';
import styles from './detail.module.less';
import classNames from 'classnames/bind';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { duotoneLight } from 'react-syntax-highlighter/dist/esm/styles/prism';

const styleCls = classNames.bind(styles);

const valueRender = (v, item) => {
  if (item.name === 'cmd' && v) {
    return (
      <SyntaxHighlighter
        className={styleCls('arg-cmd-value-wrapper')}
        language="bash"
        style={duotoneLight}
      >
        {v.replace(/;\s*/g, ';\n')}
      </SyntaxHighlighter>
    );
  } else {
    return v;
  }
};
const Inputs = ({ inputs = {} }) => {
  const resourceColumns = [
    {
      title: intl.get('type'),
      dataIndex: 'type',
    },
    {
      title: intl.get('mountPath'),
      dataIndex: 'path',
    },
  ];

  const argumentColumns = [
    {
      title: intl.get('name'),
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: intl.get('value'),
      dataIndex: 'value',
      key: 'value',
      render: valueRender,
    },
    {
      title: intl.get('description'),
      dataIndex: 'description',
      key: 'description',
    },
  ];

  return (
    <Collapse activeKey={['1', '2']}>
      <Collapse.Panel showArrow={false} header={intl.get('resources')} key="1">
        <Table
          columns={resourceColumns}
          dataSource={inputs.resources}
          pagination={false}
          rowKey="name"
        />
      </Collapse.Panel>
      <Collapse.Panel
        showArrow={false}
        header={intl.get('stage.input.arguments')}
        key="2"
      >
        <Table
          columns={argumentColumns}
          dataSource={inputs.arguments}
          pagination={false}
          rowKey="name"
        />
      </Collapse.Panel>
    </Collapse>
  );
};

Inputs.propTypes = {
  inputs: PropTypes.shape({
    resources: PropTypes.arrayOf(
      PropTypes.shape({
        name: PropTypes.string,
        path: PropTypes.string,
        type: PropTypes.string,
      })
    ).isRequired,
    arguments: PropTypes.arrayOf(
      PropTypes.shape({
        name: PropTypes.string,
        value: PropTypes.string,
      })
    ).isRequired,
  }),
};

export default Inputs;
