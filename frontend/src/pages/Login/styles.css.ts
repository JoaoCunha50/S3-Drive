import { style } from '@vanilla-extract/css'
import { Indigo, SurfaceGray, Violet } from '../../styles/colors'
import { vars } from '../../styles/theme'

export const wrapper = style({
    minHeight: '100dvh',
    display: 'grid',
    placeItems: 'center',
    padding: '0px 24px 0px 24px',
})

export const paper = style({
    width: '100%',
    maxWidth: '50%',
    backgroundColor: SurfaceGray,
    padding: '5% 10% 5% 10%',
    borderRadius: vars.radiusDefault,
})

export const header = style({
    marginBottom: '35px',
    textAlign: 'center',
})

export const title = style({
    margin: 0,
    fontWeight: 700,
    fontSize: '30px',
    background: `linear-gradient(to right, ${Indigo}, ${Violet})`,
    WebkitBackgroundClip: 'text',
    color: 'transparent',
    lineHeight: 1.2,
})

export const form = style({
    display: 'grid',
    gap: '12px',
    marginTop: '0px',
})

export const actions = style({
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginTop: '40px',
})
