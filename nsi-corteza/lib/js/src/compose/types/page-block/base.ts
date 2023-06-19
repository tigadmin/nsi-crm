import {merge} from 'lodash'
import {Apply, NoID} from '../../../cast'
import {generateUID} from '../../helpers/idgen'

interface PageBlockStyleVariants {
  [_: string]: string;
}

interface PageBlockStyleWrap {
  kind: string;
}

interface PageBlockStyleBorder {
  enabled?: boolean;
}

interface PageBlockStyle {
  variants: PageBlockStyleVariants;
  wrap?: PageBlockStyleWrap;
  border?: PageBlockStyleBorder;
}

interface PageBlockMeta {
  hidden?: boolean;
  tempID?: string;
}

export type PageBlockInput = PageBlock | Partial<PageBlock>

const defaultXYWH = [0, 0, 20, 15]

export class PageBlock {
  // blockID is auto generated by the server in order to support resource translations
  public blockID = NoID;
  public kind = ''

  public title = '';
  public description = '';

  public xywh: number[] = defaultXYWH


  public options = {}

  public meta: PageBlockMeta = {
    hidden: false,
    tempID: undefined,
  }

  public style: PageBlockStyle = {
    variants: {
      headerText: 'dark',
    },
    wrap: {
      kind: 'card',
    },
    border: {
      enabled: false,
    },
  }

  constructor (i?: PageBlockInput) {
    this.apply(i)
    this.setTempID()
  }

  apply (i?: PageBlockInput): void {
    if (!i) return

    Apply(this, i, String, 'title', 'description', 'blockID')

    if (i.xywh) {
      if (!Array.isArray(i.xywh)) {
        throw new Error('xywh must be an array')
      }

      if (i.xywh.length !== 4) {
        throw new Error('xywh must have 4 elements')
      }

      // by default, park 3x3 block in upper left corner
      this.xywh = i.xywh || defaultXYWH
    }

    if (i.options) {
      this.options = merge({}, this.options, i.options)
    }

    if (i.style) {
      this.style = merge({}, this.style, i.style)
    }

    if (i.meta) {
      this.meta = merge({}, this.meta, i.meta)
    }
  }

  // Returns Page Block configuration errors
  validate (): Array<string> {
    return []
  }

  setTempID (): void {
    this.meta.tempID = this.meta.tempID || generateUID()
  }

  clone (): PageBlockInput {
    return { ...JSON.parse(JSON.stringify(this)), blockID: NoID, meta: { tempID: '' } }
  }
}

export const Registry = new Map<string, typeof PageBlock>()
